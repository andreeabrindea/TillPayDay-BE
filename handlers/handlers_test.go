package handlers

import (
	"errors"
	"fmt"
	"internship-project3/computations"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetPayDay(t *testing.T) {
	// Create a mock HTTP request with a pay day value of 15
	req, err := http.NewRequest("GET", "/till-sallary/how-much?pay_day=15", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	//captures the response sent by an HTTP handler in the test
	rr := httptest.NewRecorder()

	// Call the GetPayDay function with the mock request and response recorder
	handler := http.HandlerFunc(GetPayDay)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	nextPayDay, _ := computations.GetNextPayDay(15, time.Now(), time.Now().Month())
	daysLeft, _ := computations.GetDaysLeft(15, time.Now(), time.Now().Month())
	// Check the response body
	expected := fmt.Sprintf(`{
  "next_pay_day": "%v",
  "days_left": %v
}`, nextPayDay.Format("January 2, 2006"), daysLeft)
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestListDates(t *testing.T) {
	// Create a mock HTTP request with a pay day value of 31
	req, err := http.NewRequest("GET", "/till-sallary/pay-day/31/list-dates", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Call the GetPayDay function with the mock request and response recorder
	handler := http.HandlerFunc(ListDates)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//var dates []NextPayDay
	// Check the response body
	//for i := time.Now(); i.Year() <= time.Now().Year(); i = i.AddDate(0, 1, 0) {
	//	output, _ := ParseNextPayDay(31, time.Now(), i, i.Month())
	//	dates = append(dates, output)
	//}
	expected := fmt.Sprintf(`{
  "next_pay_days": [
    {
      "next_pay_day": "February 28, 2023",
      "days_left": 5
    },
    {
      "next_pay_day": "March 31, 2023",
      "days_left": 36
    },
    {
      "next_pay_day": "April 28, 2023",
      "days_left": 64
    },
    {
      "next_pay_day": "May 31, 2023",
      "days_left": 97
    },
    {
      "next_pay_day": "June 30, 2023",
      "days_left": 127
    },
    {
      "next_pay_day": "July 31, 2023",
      "days_left": 158
    },
    {
      "next_pay_day": "August 31, 2023",
      "days_left": 189
    },
    {
      "next_pay_day": "September 29, 2023",
      "days_left": 218
    },
    {
      "next_pay_day": "October 31, 2023",
      "days_left": 250
    },
    {
      "next_pay_day": "November 30, 2023",
      "days_left": 280
    },
    {
      "next_pay_day": "December 29, 2023",
      "days_left": 309
    }
  ]
}`)
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
func TestParsePayDayFromURL(t *testing.T) {
	testsCases := []struct {
		name           string
		url            string
		payDayExpected int
		errExpected    error
	}{
		{
			name:           "when the payDay is 15",
			url:            "/till-sallary/pay-day/15/list-dates",
			payDayExpected: 15,
			errExpected:    nil,
		},
		{
			name:           "when payDay is a string",
			url:            "/till-sallary/pay-day/string/list-dates",
			payDayExpected: 0,
			errExpected:    errors.New("invalid URL"),
		},
		{
			name:           "when payDay is a positive integer, but not in the range 1-31",
			url:            "/till-sallary/pay-day/34/list-dates",
			payDayExpected: 0,
			errExpected:    errors.New("invalid URL"),
		},
		{
			name:           "when payDay is negative",
			url:            "/till-sallary/pay-day/-3/list-dates",
			payDayExpected: 0,
			errExpected:    errors.New("invalid URL"),
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parsePayDayFromURL(tc.url)
			if err != nil {
				if tc.errExpected == nil {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != tc.errExpected.Error() {
					t.Errorf("error mismatch (-want +got):\n%v, %v", tc.errExpected, err)
				}
			} else if got != tc.payDayExpected {
				t.Errorf("mismatch (-want +got):\n%v, %v", tc.payDayExpected, got)
			}
		})
	}
}

func TestParsePayDayFromQueryString(t *testing.T) {
	req, err := http.NewRequest("GET", "/till-sallary/how-much?pay_day=17", nil)
	if err != nil {
		t.Fatal(err)
	}
	payDay, err := ParsePayDayFromQueryString(req)
	if payDay != 17 {
		t.Errorf("Expected 17, but got %v", payDay)
	}

	req, err = http.NewRequest("GET", "/till-sallary/how-much?pay_day=string", nil)
	if err != nil {
		t.Fatal(err)
	}
	payDay, err = ParsePayDayFromQueryString(req)
	if payDay != 0 {
		t.Errorf("Expected 0, but got %v", payDay)
	}
}

func TestParseNextPayDay(t *testing.T) {
	testsCases := []struct {
		name         string
		payDay       int
		currentTime  time.Time
		markerTime   time.Time
		month        time.Month
		expectedNext NextPayDay
		expectedErr  error
	}{
		{
			name:        "when the next pay day is in the next month",
			payDay:      15,
			currentTime: time.Date(2023, time.February, 23, 0, 0, 0, 0, time.UTC),
			markerTime:  time.Date(2023, time.February, 23, 0, 0, 0, 0, time.UTC),
			month:       time.February,
			expectedNext: NextPayDay{
				NextPayDay: "March 15, 2023",
				DaysLeft:   20,
			},
			expectedErr: nil,
		},
		{
			name:        "when the next pay day is in the same month",
			payDay:      17,
			currentTime: time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC),
			markerTime:  time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC),
			month:       time.November,
			expectedNext: NextPayDay{
				NextPayDay: "November 17, 2023",
				DaysLeft:   7,
			},
			expectedErr: nil,
		},
		{
			name:         "when payDay is not in range 1-31",
			payDay:       34,
			currentTime:  time.Date(2023, time.February, 23, 0, 0, 0, 0, time.UTC),
			markerTime:   time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC),
			month:        time.February,
			expectedNext: NextPayDay{},
			expectedErr:  errors.New("pay day not in the interval 1 - 31"),
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseNextPayDay(tc.payDay, tc.currentTime, tc.markerTime, tc.month)
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("unexpected error. expected %v, but got %v", tc.expectedErr, err)
			}
			if got != tc.expectedNext {
				t.Errorf("mismatch (-want +got):\n%v, %v", tc.expectedNext, got)
			}
		})
	}
}
