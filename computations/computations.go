package computations

import "time"

func GetNextPayDay(payDay int, currentDay int, currentMonth time.Month, currentYear int) time.Time {
	date := time.Date(currentYear, currentMonth, payDay, 0, 0, 0, 0, time.UTC)

	if payDay < currentDay {
		date = date.AddDate(0, 1, 0)
	}
	if payDay > GetDaysOfCurrentMonth(currentMonth, currentYear) {
		date = time.Date(currentYear, currentMonth, GetDaysOfCurrentMonth(currentMonth, currentYear), 0, 0, 0, 0, time.UTC)
	}
	if date.Weekday() == time.Sunday {
		date = date.AddDate(0, 0, -2)
	}
	if date.Weekday() == time.Saturday {
		date = date.AddDate(0, 0, -1)
	}

	return date
}

func GetDaysLeft(payDay int, currentDay int, currentMonth time.Month, followingMonth time.Month, currentYear int) int {
	date := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.UTC)
	nextDate := GetNextPayDay(payDay, currentDay, followingMonth, currentYear)

	nextDate1 := time.Date(currentYear, nextDate.Month(), nextDate.Day(), 0, 0, 0, 0, time.UTC)
	duration := nextDate1.Sub(date)

	days := int(duration.Hours() / 24)
	return days
}

func GetDaysOfCurrentMonth(currentMonth time.Month, currentYear int) int {
	if currentMonth == time.February {
		if isLeap(currentYear) {
			return 29
		}
		return 28
	}
	if currentMonth == time.April || currentMonth == time.June || currentMonth == time.September || currentMonth == time.November {
		return 30
	}
	return 31
}

func isLeap(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 != 0 {
		return true
	}
	if year%400 == 0 {
		return true
	}
	return false
}
