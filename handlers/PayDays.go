package handlers

type PayDays struct {
	NextPayDays []NextPayDay `json:"next_pay_days"`
}

type NextPayDay struct {
	NextPayDay string `json:"next_pay_day"`
	DaysLeft   int    `json:"days_left"`
}
