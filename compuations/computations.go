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
	date := time.Date(currentYear, currentMonth, payDay, 0, 0, 0, 0, time.UTC)
	if isSaturday(payDay, currentMonth, currentYear) == true {
		date = date.AddDate(0, 0, -1) // add -1 days to move to Friday
	} else if isSunday(payDay, currentMonth, currentYear) == true {
		date = date.AddDate(0, 0, -2) // add -2 to move to Frida
	}
	return date.Day()
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
	if date.Weekday() == time.Saturday {
		return true
	}
	return false
}

func GetDaysLeft(payDay int, currentDay int, currentMonth time.Month, followingMonth time.Month, currentYear int) int {
	date := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.UTC)
	nextDate := time.Date(currentYear, followingMonth, GetNextPayDay(payDay, currentMonth, currentYear), 0, 0, 0, 0, time.UTC)
	duration := nextDate.Sub(date)

	//the number of days between the two dates
	days := int(duration.Hours() / 24)
	return days
}

func GetMonthOfSalary(payDay int, currentDay int, currentMonth time.Month) int {
	if currentDay >= payDay {
		return int(currentMonth + 1)
	}
	return int(currentMonth)
}
