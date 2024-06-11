package main

import "fmt"

func getInputValues(text string) float64 {
	fmt.Print(text)
	var inputValue float64
	fmt.Scan(&inputValue)
	return inputValue
}
