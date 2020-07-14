package main

import "testing"

func TestGetMonthNameJanuary(t *testing.T) {
	monthName, err := getMonthName(1)
	if err != nil {
		t.Errorf("getMonthName(1) got an error: %s\n", err)
	}

	if monthName != "January" {
		t.Errorf("getMonthName(1) exected January, got %s\n", monthName)
	}
}
func TestGetMonthNameInvalidMonth(t *testing.T) {
	_, err := getMonthName(13)
	if err == nil {
		t.Errorf("getMonthName(13) returned nil error\n")
	}

	_, err = getMonthName(0)
	if err == nil {
		t.Errorf("getMonthName(0) returned nil error\n")
	}
}
