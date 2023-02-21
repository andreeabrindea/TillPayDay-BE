package main

import (
	"encoding/json"
	"fmt"
	"internship-project3/compuations"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetPayDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pay_day := r.URL.Query().Get("pay_day")
	fmt.Println("day =", pay_day)
	payDay, err := strconv.Atoi(pay_day)
	if err != nil {
		_, err = w.Write([]byte("unexpected error"))
		if err != nil {
			return
		}
		return
	}
	err = compuations.ValidateDay(payDay)
	if err != nil {
		_, err := w.Write([]byte("Wrong date"))
		if err != nil {
			return
		}
		return
	}
	date := Data{Day: compuations.GetNextPayDay(payDay, time.Now().Month(), time.Now().Year()), Month: compuations.GetMonthOfSalary(payDay, time.Now().Day(), time.Now().Month())}
	output := NextPayDate{NextDate: date, DaysLeft: compuations.GetDaysLeft(payDay, time.Now().Day(), time.Now().Month(), time.Now().Year())}
	msg, _ := json.MarshalIndent(output, "", "")
	_, _ = w.Write(msg)

}

func ListDates(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	payDayParam := q.Get("pay_day")
	if payDayParam == "" {
		http.Error(w, "Missing pay_day parameter", http.StatusBadRequest)
		return
	}

	payDayParts := strings.Split(payDayParam, "/")
	payDay, err := strconv.Atoi(payDayParts[0])
	if err != nil {
		http.Error(w, "Invalid pay_day parameter", http.StatusBadRequest)
		return
	}

	// Calculate next pay days
	now := time.Now()
	year := time.Now().Year()
	var nextPayDays []string
	for i := 1; i <= 12; i++ {
		t := time.Date(year, time.Month(i), payDay, 0, 0, 0, 0, time.UTC)
		if t.Before(now) {
			t = t.AddDate(1, 0, 0)
		}
		nextPayDays = append(nextPayDays, t.Format("January 2, 2006"))
	}

	// Write response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Next pay days for day %d:\n\n", payDay)
	for _, date := range nextPayDays {
		fmt.Fprintln(w, date)
	}
}
func main() {
	http.HandleFunc(
		"/till-salary/how-much",
		GetPayDay,
	)
	http.HandleFunc("/till-salary/pay-day/", ListDates)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
