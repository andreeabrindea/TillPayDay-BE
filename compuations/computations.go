package compuations

import (
	"errors"
	"time"
)

func ValidateDay(payDay int) error {
	if payDay > 31 || payDay < 1 {
		return errors.New("invalid payDay")
	}
	return nil
}
func GetNextPayDay(payDay int, currentMonth time.Month, currentYear int) int {
	nextPayDate := payDay
	if isSaturday(payDay, currentMonth, currentYear) == true {
		nextPayDate = payDay - 1
	}
	if isSunday(payDay, currentMonth, currentYear) == true {
		nextPayDate = payDay - 2
	}
	if payDay > GetDaysOfCurrentMonth(currentMonth, currentYear) {
		nextPayDate = GetDaysOfCurrentMonth(currentMonth, currentYear)
	}
	return nextPayDate
}

func isSunday(payDay int, currentMonth time.Month, currentYear int) bool {
	// Create a time object for the given payDay
	date := time.Date(currentYear, currentMonth, payDay, 0, 0, 0, 0, time.UTC)

	if date.Weekday() == time.Sunday {
		return true
	}
	return false
}
func isSaturday(payDay int, currentMonth time.Month, currentYear int) bool {
	date := time.Date(currentYear, currentMonth, payDay, 0, 0, 0, 0, time.UTC)

	// Check if the weekday is Saturday or Sunday (time.Saturday and time.Sunday, respectively, in Go's weekday numbering)
	if date.Weekday() == time.Saturday {
		return true
	}
	return false
}

func GetDaysLeft(payDay int, currentDay int, currentMonth time.Month, currentYear int) int {
	daysLeft := GetNextPayDay(payDay, currentMonth, currentYear) - currentDay

	if payDay > GetDaysOfCurrentMonth(currentMonth, currentYear) {
		daysLeft = GetDaysOfCurrentMonth(currentMonth, currentYear) - currentDay

	}

	if GetNextPayDay(payDay, currentMonth, currentYear) < currentDay {
		daysLeft = GetDaysOfCurrentMonth(currentMonth, currentYear) - currentDay + GetNextPayDay(payDay, currentMonth, currentYear)
	}
	return daysLeft
}

func GetMonthOfSalary(payDay int, currentDay int, currentMonth time.Month) int {
	if payDay < currentDay {
		return int(currentMonth + 1)
	}
	return int(currentMonth)
}

func GetDaysOfCurrentMonth(currentMonth time.Month, currentYear int) int {
	if currentMonth == 2 {
		if currentYear/4 == 0 {
			return 29
		} else {
			return 28
		}
	}

	if currentMonth == 1 || currentMonth == 3 || currentMonth == 5 || currentMonth == 7 || currentMonth == 8 || currentMonth == 10 || currentMonth == 12 {
		return 31
	}
	return 30
}
