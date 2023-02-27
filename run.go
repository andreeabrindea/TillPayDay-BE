package main

import (
	"fmt"
	"internship-project3/computations"
	"internship-project3/handlers"
	"net/http"
)

func main() {
	_, err := computations.GetRomanianHolidays("postgres://xvyctfje:5yGXTCPQKkKJe0rjuvsJtFOQF7BiOBJp@mouse.db.elephantsql.com/xvyctfje")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.HandleFunc(
		"/till-sallary/how-much",
		handlers.GetPayDay,
	)
	http.HandleFunc("/till-sallary/pay-day/", handlers.ListDates)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
