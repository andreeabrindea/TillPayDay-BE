package computations

import (
	"testing"
	"time"
)

func TestIsSaturday(t *testing.T) {
	date := time.Date(2022, 12, 24, 0, 0, 0, 0, time.UTC)
	if isSaturday(24, 12, 2022) != true {
		t.Errorf("Expected %v to be on saturday", date)
	}

	date = time.Date(2023, 02, 21, 0, 0, 0, 0, time.UTC)
	if isSaturday(21, 02, 2023) != false {
		t.Errorf("Expected %v to be no on weekend", date)
	}

}

func TestIsSunday(t *testing.T) {
	date := time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC)
	if isSunday(25, 12, 2022) != true {
		t.Errorf("Expected %v to be on saturday", date)
	}

	date = time.Date(2023, 02, 21, 0, 0, 0, 0, time.UTC)
	if isSunday(21, 02, 2023) != false {
		t.Errorf("Expected %v to be no on weekend", date)
	}
}
func TestGetDaysLeft(t *testing.T) {

}
