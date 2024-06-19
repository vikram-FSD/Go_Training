package main

import (
	"errors"
	"fmt"
)

func getInputValues(text string) (float64, error) {
	fmt.Print(text)
	var inputValue float64
	fmt.Scan(&inputValue)
	if inputValue < 0 || inputValue == 0 {
		return 0, errors.New("invalid input Negative values not allowed please enter more than 0")
	} else {
		return inputValue, nil
	}
}
