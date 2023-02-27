package main

import (
	"fmt"
	"internship-project3/computations"
	"internship-project3/handlers"
	"net/http"
)

func main() {
	//make this run just one time, and not at every call, so it doesn't affect the performance too much
	//the output is stored in a global variable (holidays)
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
		fmt.Println(err)
		return
	}
}
