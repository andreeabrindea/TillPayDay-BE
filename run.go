package main

import (
	"fmt"
	"internship-project3/handlers"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/till-sallary/how-much",
		handlers.GetPayDay,
	)
	http.HandleFunc("/till-sallary/pay-day/", handlers.ListDates)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
