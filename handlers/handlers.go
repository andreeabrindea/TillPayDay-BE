package handlers

import (
	"encoding/json"
	"fmt"
	"internship-project3/computations"
	"time"

	"net/http"
	"strconv"
	"strings"
)

func GetPayDay(w http.ResponseWriter, r *http.Request) {
	pay_day := r.URL.Query().Get("pay_day")
	payDay, err := strconv.Atoi(pay_day)
	if err != nil {
		_, err = w.Write([]byte("Argument couldn't be casted to int"))
		if err != nil {
			return
		}
		return
	}
	if payDay > 31 || payDay < 1 {
		_, err = w.Write([]byte("Invalid argument"))
		if err != nil {
			return
		}
		return
	}
	now := time.Now()
	nextPayDay := computations.GetNextPayDay(payDay, now.Day(), now.Month(), now.Year())
	daysLeft := computations.GetDaysLeft(payDay, now.Day(), now.Month(), nextPayDay.Month(), now.Year())
	output := NextPayDay{NextPayDay: nextPayDay.Format("January 2, 2006"), DaysLeft: daysLeft}
	u, err := json.MarshalIndent(output, "", "")
	_, err = w.Write(u)
	if err != nil {
		_, err = w.Write([]byte("Couldn't write response"))
		if err != nil {
			return
		}
		return
	}
}

func ListDates(w http.ResponseWriter, r *http.Request) {
	// Extract the pay day from the URL path
	parts := strings.Split(r.URL.Path, "/")

	if URLContains(parts, "list-dates") != true {
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
	fmt.Println(payDay)
	var dates []NextPayDay
	for i := time.Now().Month(); i <= 12; i++ {
		nextPayDay := computations.GetNextPayDay(payDay, time.Now().Day(), i, time.Now().Year())
		nextPayDayFormatted := NextPayDay{NextPayDay: nextPayDay.Format("January 2, 2006"), DaysLeft: computations.GetDaysLeft(payDay, time.Now().Day(), time.Now().Month(), i, time.Now().Year())}
		dates = append(dates, nextPayDayFormatted)
		if nextPayDay.Month() == 12 {
			break
		}
	}
	output := PayDays{dates}
	u, err := json.MarshalIndent(output, "", "")
	_, err = w.Write(u)
	if err != nil {
		_, err = w.Write([]byte("Couldn't write response"))
		if err != nil {
			return
		}
		return
	}
}
func URLContains(url []string, substring string) bool {
	ok := false
	for i := range url {
		if strings.Contains(url[i], substring) == true {
			ok = true
		}
	}
	return ok
}
