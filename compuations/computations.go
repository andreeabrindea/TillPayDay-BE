package compuations

import (
	"errors"
	"time"
)

func ValidateDay(day int) error {
	if day > 31 || day < 1 {
		return errors.New("invalid day")
	}
	return nil
}
func CalculateNextPayDay(day int) int {
	nextPayDate := day
	if isWeekend(day) == "Saturday" {
		nextPayDate = day - 1
	}
	if isWeekend(day) == "Sunday" {
		nextPayDate = day - 2
	}
	return nextPayDate
}

func isWeekend(day int) string {
	year, month, _ := time.Now().Date()

	// Create a time object for the given day
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	// Check if the weekday is Saturday or Sunday (time.Saturday and time.Sunday, respectively, in Go's weekday numbering)
	if date.Weekday() == time.Saturday {
		return "Saturday"
	}
	if date.Weekday() == time.Sunday {
		return "Sunday"
	}
	return ""
}

func CalculateDaysLeft(day int) int {
	daysLeft := CalculateNextPayDay(day) - time.Now().Day()

	if day > GetDaysOfCurrentMonth() {
		daysLeft = GetDaysOfCurrentMonth() - time.Now().Day()

	}

	if CalculateNextPayDay(day) < time.Now().Day() {
		daysLeft = GetDaysOfCurrentMonth() - time.Now().Day() + CalculateNextPayDay(day)
	}
	return daysLeft
}

func GetMonthOfSalary(day int) int {
	if day < time.Now().Day() {
		return int(time.Now().Month() + 1)
	}
	return int(time.Now().Month())
}

func GetDaysOfCurrentMonth() int {
	if time.Now().Month() == 2 {
		if time.Now().Year()/4 == 0 {
			return 29
		} else {
			return 28
		}
	}

	if time.Now().Month() == 1 || time.Now().Month() == 3 || time.Now().Month() == 5 || time.Now().Month() == 7 || time.Now().Month() == 8 || time.Now().Month() == 10 || time.Now().Month() == 12 {
		return 31
	}
	return 30
}
