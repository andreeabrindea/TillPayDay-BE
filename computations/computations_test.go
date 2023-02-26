package computations

import (
	"errors"
	"testing"
	"time"
)

func TestGetNextPayDay(t *testing.T) {
	testsCases := []struct {
		name            string
		payDay          int
		currentTime     time.Time
		markerTime      time.Time
		month           time.Month
		expectedNextDay time.Time
		expectedErr     error
	}{
		{
			name:            "when the next pay day is in the next month",
			payDay:          15,
			currentTime:     time.Date(2023, time.February, 23, 0, 0, 0, 0, time.Local),
			month:           time.February,
			expectedNextDay: time.Date(2023, time.March, 15, 0, 0, 0, 0, time.Local),
			expectedErr:     nil,
		},
		{
			name:            "when the next pay day is in the same month",
			payDay:          17,
			currentTime:     time.Date(2023, time.November, 10, 0, 0, 0, 0, time.Local),
			month:           time.November,
			expectedNextDay: time.Date(2023, time.November, 17, 0, 0, 0, 0, time.Local),
			expectedErr:     nil,
		},
		{
			name:            "when payDay is not in range 1-31",
			payDay:          34,
			currentTime:     time.Date(2023, time.February, 23, 0, 0, 0, 0, time.Local),
			month:           time.February,
			expectedNextDay: time.Time{},
			expectedErr:     errors.New("pay day not in the interval 1 - 31"),
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetNextPayDay(tc.payDay, tc.currentTime, tc.month)
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("unexpected error. expected %v, but got %v", tc.expectedErr, err)
			}
			if got != tc.expectedNextDay {
				t.Errorf("mismatch (-want +got):\n%v, %v", tc.expectedNextDay, got)
			}
		})
	}
}
func TestGetDaysLeft(t *testing.T) {
	testsCases := []struct {
		name             string
		payDay           int
		currentTime      time.Time
		markMonth        time.Month
		expectedNoOfDays int
		expectedErr      error
	}{
		{
			name:             "when the next pay day is in the next month",
			payDay:           15,
			currentTime:      time.Date(2023, time.February, 23, 0, 0, 0, 0, time.Local),
			markMonth:        time.February,
			expectedNoOfDays: 20,
			expectedErr:      nil,
		},
		{
			name:             "when the next pay day is in the same month",
			payDay:           17,
			currentTime:      time.Date(2023, time.November, 10, 0, 0, 0, 0, time.Local),
			markMonth:        time.November,
			expectedNoOfDays: 7,
			expectedErr:      nil,
		},
		{
			name:             "when payDay is not in range 1-31",
			payDay:           34,
			currentTime:      time.Date(2023, time.February, 23, 0, 0, 0, 0, time.Local),
			markMonth:        time.February,
			expectedNoOfDays: 0,
			expectedErr:      errors.New("pay day not in the interval 1 - 31"),
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetDaysLeft(tc.payDay, tc.currentTime, tc.markMonth)
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("unexpected error. expected %v, but got %v", tc.expectedErr, err)
			}
			if got != tc.expectedNoOfDays {
				t.Errorf("mismatch (-want +got):\n%v, %v", tc.expectedNoOfDays, got)
			}
		})
	}
}

func TestGetDaysOfCurrentMonth(t *testing.T) {
	testsCases := []struct {
		name             string
		month            time.Month
		currentTime      time.Time
		expectedNoOfDays int
	}{
		{
			name:             "when the month has 28 days",
			month:            time.February,
			currentTime:      time.Now(),
			expectedNoOfDays: 28,
		},
		{
			name:             "when the month has 29 days",
			month:            time.February,
			currentTime:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			expectedNoOfDays: 29,
		},
		{
			name:             "when the month has 30 days",
			month:            time.April,
			currentTime:      time.Now(),
			expectedNoOfDays: 30,
		},
		{
			name:             "when the month has 31 days",
			month:            time.March,
			currentTime:      time.Now(),
			expectedNoOfDays: 31,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := GetDaysOfCurrentMonth(tc.month, tc.currentTime)
			if got != tc.expectedNoOfDays {
				t.Errorf("mismatch (-want +got):\n%v, %v", tc.expectedNoOfDays, got)
			}
		})
	}
}

func TestIsLeap(t *testing.T) {
	testsCases := []struct {
		name          string
		year          int
		expected      bool
		expectedError error
	}{
		{
			name:          "when year is negative",
			year:          -200,
			expected:      false,
			expectedError: errors.New("unexpected year"),
		},

		{
			name:          "when year is valid and leap",
			year:          2024,
			expected:      true,
			expectedError: nil,
		},
		{
			name:          "when year is valid and non-leap",
			year:          2023,
			expected:      false,
			expectedError: nil,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IsLeap(tc.year)
			if got != tc.expected && err != tc.expectedError {
				t.Errorf("mismatch (-wanted value,error:, +got value, error: ):\n%v,%v, %v, %v", tc.expected, tc.expectedError, got, err)
			}
		})
	}
}
func TestIsPublicHoliday(t *testing.T) {
	testsCases := []struct {
		name        string
		payDay      int
		month       time.Month
		currentTime time.Time
		expected    bool
	}{
		{
			name:        "when a public holiday that falls on the given day and month",
			payDay:      1,
			month:       time.January,
			currentTime: time.Now(),
			expected:    true,
		},
		{
			name:        "when the given day is not a public holiday",
			payDay:      17,
			month:       time.November,
			currentTime: time.Now(),
			expected:    false,
		},
	}
	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPublicHoliday(tc.payDay, tc.month, tc.currentTime)
			if got != tc.expected {
				t.Errorf("mismatch (-wanted, +got):\n%v,%v", tc.expected, got)
			}
		})
	}

}
