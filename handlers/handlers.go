package handlers

import (
	"encoding/json"
	"errors"
	"internship-project3/computations"
	"regexp"
	"time"

	"net/http"
	"strconv"
	"strings"
)

// ParsePayDay extracts the pay day value from the URL query parameters of the request, converts it to an integer using the strconv.Atoi function.
func ParsePayDay(r *http.Request) (int, error) {
	payDayStr := r.URL.Query().Get("pay_day")
	payDay, err := strconv.Atoi(payDayStr)
	if err != nil {
		return 0, errors.New("pay day should be an integer")
	}
	return payDay, nil
}

// ParseNextPayDay acts like a coordinator between ParsePayDay and GetPayDay.
// It calculates the next pay and returns a NextPayDay struct that contains the formatted next pay day and the number of days left.
// Note that currentTime and markerTime could be the same if we want to calculate the days left from now to the next pay day,
// but if we want to calculate the days from now to the fifth salary day, then currentTime will be time.Now and markerTime the fifth month
func ParseNextPayDay(payDay int, currentTime time.Time, markerTime time.Time, month time.Month) (NextPayDay, error) {
	nextPayDay, err := computations.GetNextPayDay(payDay, markerTime, month)
	if err != nil {
		return NextPayDay{}, err
	}
	daysLeft, err := computations.GetDaysLeft(payDay, currentTime, month)
	if err != nil {
		return NextPayDay{}, err
	}
	output := NextPayDay{
		NextPayDay: nextPayDay.Format("January 2, 2006"),
		DaysLeft:   daysLeft,
	}
	return output, nil
}

// GetPayDay handles an HTTP request and returns a JSON response containing information about the next pay day
func GetPayDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	payDay, err := ParsePayDay(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := ParseNextPayDay(payDay, time.Now(), time.Now(), time.Now().Month())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "  ")
	err = jsonEncoder.Encode(output)

	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

// ListDates handles an HTTP request and returns a JSON response containing information about the next pay days
func ListDates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	payDay, err := parsePayDayFromUglyURL(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dates []NextPayDay

	// i iterates over each month starting from the current month until the end of the current year.
	for i := time.Now(); i.Year() <= time.Now().Year(); i = i.AddDate(0, 1, 0) {
		output, err := ParseNextPayDay(payDay, time.Now(), i, i.Month())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dates = append(dates, output)
	}
	w.Header().Set("Content-Type", "application/json")
	nextDates := PayDays{NextPayDays: dates}
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "  ")
	err = jsonEncoder.Encode(nextDates)

	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

}
func parsePayDayFromUglyURL(urlPath string) (int, error) {
	parts := strings.Split(urlPath, "/")
	regExp := regexp.MustCompile("^/till-salary/pay-day/(0?[1-9]|[1-2][0-9]|3[0-1])/list-dates$")

	if regExp.MatchString(urlPath) == false {
		return 0, errors.New("invalid URL")
	}

	payDay, err := strconv.Atoi(parts[3])
	if err != nil {
		return 0, errors.New("given value is not an integer")
	}
	return payDay, nil
}
