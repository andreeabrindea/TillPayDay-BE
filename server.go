package main

import (
	"encoding/json"
	"fmt"
	"internship-project3/compuations"
	"net/http"
	"strconv"
	"strings"
)

func GetPayDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pay_day := r.URL.Query().Get("pay_day")
	fmt.Println("day =", pay_day)
	payDay, _ := strconv.Atoi(pay_day)

	err := compuations.ValidateDay(payDay)
	if err != nil {
		_, err := w.Write([]byte("Wrong date"))
		if err != nil {
			return
		}
		return
	}
	date := Data{Day: compuations.CalculateNextPayDay(payDay), Month: compuations.GetMonthOfSalary(payDay)}
	output := NextPayDate{NextDate: date, DaysLeft: compuations.CalculateDaysLeft(payDay)}
	msg, _ := json.MarshalIndent(output, "", "")
	_, _ = w.Write(msg)

}

func ListDates(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	payDay, err := strconv.Atoi(parts[3])

	dates := make([]string, 0)
	for i := 0; i < 5; i++ {
		dates = append(dates, fmt.Sprintf("2023-02-%02d", payDay+i))
	}

	// return the list of dates as a JSON response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(fmt.Sprintf(`{"pay_day": %d, "dates": %v}`, payDay, dates)))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc(
		"/till-salary/how-much",
		GetPayDay,
	)
	http.HandleFunc(
		"/till-salary/pay-day",
		ListDates)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
