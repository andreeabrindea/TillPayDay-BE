package main

import (
	"internship-project3/handlers"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/till-salary/how-much",
		handlers.GetPayDay,
	)
	http.HandleFunc("/till-salary/pay-day/", handlers.ListDates)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
