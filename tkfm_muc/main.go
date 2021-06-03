package main

import "fmt"

func main() {
	var days int
	var demonites int = 0

	fmt.Println("Please input you want to make up days:")
	_, e := fmt.Scanln(&days)
	if e != nil {
		fmt.Println("Error: Days must be set to interger.")
		return
	}
	if days <= 0 {
		fmt.Println("Error: Make up days must be greater than '0'")
		return
	}

	for i := 0; i < days; i++ {
		demonites += (50 + 5*i)
	}

	fmt.Printf("Ok, you shall spend '%d' Demonties.\n", demonites)
}
