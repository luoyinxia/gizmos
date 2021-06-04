package main

import (
	"fmt"
	"time"
)

func main() {
	var days int
	var demonites int = 0

	// get user input data
	fmt.Println("Please input you want to make up days:")
	_, e := fmt.Scanln(&days)
	// error handlers
	ok := daysChecker(days, e)
	if !ok {
		return
	}

	// calculate
	for i := 0; i < days; i++ {
		demonites += (50 + 5*i)
	}

	// show result
	fmt.Printf("Ok, you shall spend '%d' Demonties.\n", demonites)
}

func daysChecker(days int, e error) bool {
	// default error handler
	if e != nil {
		fmt.Println("Error: Days must be set to integer.")
		return false
	}
	// input days checker
	if days <= 0 {
		fmt.Println("Error: Make up days must be greater than '0'.")
		return false
	}
	// compare input days with days of current month
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month().String()
	switch currentMonth {
	case "January", "March", "May", "July", "August", "October", "December":
		if days > 31 {
			fmt.Println("Error: The current month only have 31 days.")
			return false
		}
	case "April", "June", "September", "November":
		if days > 30 {
			fmt.Println("Error: The current month only have 30 days.")
			return false
		}
	case "February":
		// days of february(common year)
		var daysOfFeb int = 28
		// days of february(leap year)
		if (currentYear%100 == 0) && (currentYear%400 == 0) {
			daysOfFeb = 29
		} else if (currentYear%100 != 0) && (currentYear%4 == 0) {
			daysOfFeb = 29
		}

		if days > daysOfFeb {
			fmt.Printf("Error: The current month only have %d days.\n", daysOfFeb)
			return false
		}
	}
	return true
}
