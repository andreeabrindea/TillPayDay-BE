package computations

import (
	"errors"
	"log"
	"time"
)

// GetNextPayDay calculates the next pay day for the given month and year.
// payDay = the day of the month on which the pay day falls
// currentTime = the current date and time
// month = the month for which the pay day is being calculated
func GetNextPayDay(payDay int, currentTime time.Time, month time.Month) (time.Time, error) {
	if payDay < 1 || payDay > 31 {
		return time.Time{}, errors.New("pay day not in the interval 1 - 31")
	}
	// the date will consist of the pay day and the month given
	date := time.Date(currentTime.Year(), month, payDay, 0, 0, 0, 0, time.Local)

	// check if the pay day is bigger than the number of days of the month ex. payDay=31 and month=2
	// then the payDay will be on the last day of the month => payDay=28/29
	noOfDays, _ := GetDaysOfCurrentMonth(month, currentTime)
	if noOfDays < payDay {
		date = time.Date(currentTime.Year(), month, noOfDays, 0, 0, 0, 0, time.Local)
	}

	// check if the pay day of the month has passed if yes => we add a month
	if payDay < currentTime.Day() {
		date = date.AddDate(0, 1, 0)
	}
	if date.Weekday() == time.Sunday {
		date = date.AddDate(0, 0, 1)
	}
	if date.Weekday() == time.Saturday {
		date = date.AddDate(0, 0, 2)
	}
	for {
		if IsHoliday(date) == true {
			date = date.AddDate(0, 0, 1)
		}
		if date.Weekday() == time.Sunday {
			date = date.AddDate(0, 0, 1)
		}
		if date.Weekday() == time.Saturday {
			date = date.AddDate(0, 0, 2)
		}
		if IsHoliday(date) == false {
			break
		}
	}
	return date, nil
}

// GetDaysLeft calculates the number of days between 2 dates.
func GetDaysLeft(payDay int, currentTime time.Time, markMonth time.Month) (int, error) {
	// date is the current time
	date := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)

	// Calculate next pay day for the given month and year
	nextDate, err := GetNextPayDay(payDay, currentTime, markMonth)
	if err != nil {
		return 0, err
	}

	// If next pay day is before the current date, calculate next pay day for the following year
	if nextDate.Before(date) {
		nextDate, err = GetNextPayDay(payDay, time.Date(currentTime.Year()+1, markMonth, 1, 0, 0, 0, 0, time.Local), markMonth)
		if err != nil {
			return 0, err
		}
	}

	nextDate = time.Date(nextDate.Year(), nextDate.Month(), nextDate.Day(), 0, 0, 0, 0, time.Local)
	duration := nextDate.Sub(date)

	days := int(duration.Hours() / 24)
	return days, nil
}

// GetDaysOfCurrentMonth returns the number of days of the given month and time
// it takes into consideration if the year leap or not
func GetDaysOfCurrentMonth(month time.Month, currentTime time.Time) (int, error) {
	if month == 0 {
		return 0, errors.New("unexpected month")
	}
	if month == time.February {
		isLeap, err := IsLeap(currentTime.Year())
		if err != nil {
			log.Print(err)
			return 0, err
		}
		if isLeap == true {
			return 29, nil
		}
		return 28, nil
	}
	if month == time.April || month == time.June || month == time.September || month == time.November {
		return 30, nil
	}
	return 31, nil
}

// IsLeap checks if the given year is leap or not
func IsLeap(year int) (bool, error) {
	if year < 0 || year > 9999 { //9999 is the biggest year supported by time.Time.Year()
		return false, errors.New("unexpected year")
	}
	if year%4 != 0 {
		return false, nil
	}
	if year%100 != 0 {
		return true, nil
	}
	if year%400 == 0 {
		return true, nil
	}
	return false, nil
}