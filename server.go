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
	pay_day := r.URL.Query().Get("pay_day")
	payDay, err := strconv.Atoi(pay_day)
	if err != nil {
		_, err = w.Write([]byte("argument couldn't be casted to int"))
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

	w.Header().Set("Content-Type", "application/json")
	date := Data{Day: compuations.GetNextPayDay(payDay, time.Now().Month(), time.Now().Year()), Month: compuations.GetMonthOfSalary(payDay, time.Now().Day(), time.Now().Month())}
	output := NextPayDate{NextDate: date, DaysLeft: compuations.GetDaysLeft(compuations.GetNextPayDay(payDay, time.Now().Month(), time.Now().Year()), time.Now().Day(), time.Now().Month(), time.Month(compuations.GetMonthOfSalary(payDay, time.Now().Day(), time.Now().Month())), time.Now().Year())}
	msg, _ := json.MarshalIndent(output, "", "")
	_, _ = w.Write(msg)

}

func ListDates(w http.ResponseWriter, r *http.Request) {
	// Extract the pay day from the URL path
	parts := strings.Split(r.URL.Path, "/")
	fmt.Println(parts)
	ok := 0
	for i := range parts {
		if strings.Contains(parts[i], "list-dates") == true {
			ok = 1
		}
	}
	if ok != 1 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	if len(parts) != 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	payDay, err := strconv.Atoi(parts[3])
	if err != nil {
		http.Error(w, "Invalid pay day", http.StatusBadRequest)
		return
	}

	// Calculate next pay days
	err = compuations.ValidateDay(payDay)
	if err != nil {
		_, err := w.Write([]byte("wrong date"))
		if err != nil {
			return
		}
		return
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	var dates []NextPayDate
	currentDay := time.Now().Day()
	for i := time.Now().Month(); i < 12; i++ {
		date := Data{Day: compuations.GetNextPayDay(payDay, i+1, time.Now().Year()), Month: compuations.GetMonthOfSalary(payDay, currentDay, i)}
		nextPayDay := NextPayDate{NextDate: date, DaysLeft: compuations.GetDaysLeft(compuations.GetNextPayDay(payDay, i+1, time.Now().Year()), time.Now().Day(), time.Now().Month(), i+1, time.Now().Year())}
		dates = append(dates, nextPayDay)
	}
	output := PayDay{Dates: dates}
	msg, _ := json.MarshalIndent(output, "", "")
	_, _ = w.Write(msg)
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
