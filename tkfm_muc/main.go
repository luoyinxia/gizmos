package main

import "fmt"

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
	if e != nil {
		fmt.Println("Error: Days must be set to integer.")
		return false
	}
	if days <= 0 {
		fmt.Println("Error: Make up days must be greater than '0'")
		return false
	}
	return true
}
