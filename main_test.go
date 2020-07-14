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
func TestGetMonthNameFebruary(t *testing.T) {
	monthName, err := getMonthName(2)
	if err != nil {
		t.Errorf("getMonthName(2) got an error: %s\n", err)
	}

	if monthName != "February" {
		t.Errorf("getMonthName(2) exected February, got %s\n", monthName)
	}
}
func TestGetMonthNameMarch(t *testing.T) {
	monthName, err := getMonthName(3)
	if err != nil {
		t.Errorf("getMonthName(3) got an error: %s\n", err)
	}

	if monthName != "March" {
		t.Errorf("getMonthName(3) exected March, got %s\n", monthName)
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
