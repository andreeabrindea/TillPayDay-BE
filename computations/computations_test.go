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
			currentTime:     time.Date(2023, time.February, 23, 0, 0, 0, 0, time.UTC),
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
func TestParseNextPayDay(t *testing.T) {
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
