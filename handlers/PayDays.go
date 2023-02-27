package handlers

type PayDays struct {
	NextPayDays []NextPayDay `json:"next_pay_days"`
}

type NextPayDay struct {
	NextDay  string `json:"next_pay_day"`
	DaysLeft int    `json:"days_left"`
}
