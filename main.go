package main

import (
	"errors"
	"fmt"
	"rsc.io/quote"
)

func getMonthName(month int) (name string, err error) {
	if month < 1 || month > 12 {
		return "", errors.New("invalid month")
	}

	months := [12]string{"January", "Februaryy", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	return months[month-1], nil
}

func main() {
	fmt.Println(quote.Hello())

	month := 1
	monthName, err := getMonthName(month)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("The month with number %d is %s.\n", month, monthName)
	}
}
