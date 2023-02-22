package main

type PayDay struct {
	Dates []NextPayDate `json:"dates"`
}

type NextPayDate struct {
	NextDate Data `json:"next_date"`
	DaysLeft int  `json:"days_left"`
}

type Data struct {
	Day   int `json:"day"`
	Month int `json:"month"`
}
