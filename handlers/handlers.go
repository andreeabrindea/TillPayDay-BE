package handlers

import (
	"encoding/json"
	"errors"
	"internship-project3/computations"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParsePayDayFromQueryString extracts the pay day value from the URL query parameters of the request, converts it to an integer using the strconv.Atoi function.
func ParsePayDayFromQueryString(r *http.Request) (int, error) {
	payDayStr := r.URL.Query().Get("pay_day")
	payDay, err := strconv.Atoi(payDayStr)
	if err != nil {
		return 0, errors.New("pay day should be an integer")
	}
	if payDay < 1 || payDay > 31 {
		return 0, errors.New("pay day should be in range 1-31")
	}
	return payDay, nil
}

// ParseNextPayDay acts like a coordinator between ParsePayDayFromQueryString and GetPayDay.
// It calculates the next pay and returns a NextPayDay struct that contains the formatted next pay day and the number of days left.
// Note that currentTime and markerTime could be the same if we want to calculate the days left from now to the next pay day,
// but if we want to calculate the days from now to the fifth salary day, then currentTime will be time.Now and markerTime the fifth month
func ParseNextPayDay(payDay int, currentTime time.Time, markerTime time.Time, month time.Month) (NextPayDay, error) {
	nextPayDay, err := computations.GetNextPayDay(payDay, markerTime, month)
	if nextPayDay.Year() != currentTime.Year() {
		return NextPayDay{}, nil
	}
	if err != nil {
		return NextPayDay{}, err
	}
	daysLeft, err := computations.GetDaysLeft(payDay, currentTime, month)
	if err != nil {
		return NextPayDay{}, err
	}
	output := NextPayDay{
		NextDay:  nextPayDay.Format("January 2, 2006"),
		DaysLeft: daysLeft,
	}
	return output, nil
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// GetPayDay handles an HTTP request and returns a JSON response containing information about the next pay day
func GetPayDay(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	payDay, err := ParsePayDayFromQueryString(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	output, err := ParseNextPayDay(payDay, time.Now(), time.Now(), time.Now().Month())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	u, _ := json.MarshalIndent(output, "", "  ")
	_, err = w.Write(u)
	if err != nil {
		return
	}
}

// ListDates handles an HTTP request and returns a JSON response containing information about the next pay days
func ListDates(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	payDay, err := parsePayDayFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var dates []NextPayDay
	currentTime := time.Now()

	// i iterates over each month starting from the current month until the end of the current year.
	for i := time.Now(); i.Year() == time.Now().Year(); i = i.AddDate(0, 1, 0) {
		output, err := ParseNextPayDay(payDay, currentTime, i, i.Month())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if output != (NextPayDay{}) {
			dates = append(dates, output)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	nextDates := PayDays{NextPayDays: dates}
	next, _ := json.MarshalIndent(nextDates, "", "")
	_, err = w.Write(next)
	if err != nil {
		return
	}
}
func parsePayDayFromURL(urlPath string) (int, error) {
	parts := strings.Split(urlPath, "/")
	regExp := regexp.MustCompile("^/till-sallary/pay-day/(0?[1-9]|[1-2][0-9]|3[0-1])/list-dates$")

	if regExp.MatchString(urlPath) == false {
		return 0, errors.New("invalid URL! Try an integer number in range 1-31")
	}

	payDay, err := strconv.Atoi(parts[3])
	if err != nil {
		return 0, errors.New("given value is not an integer")
	}
	return payDay, nil
}
