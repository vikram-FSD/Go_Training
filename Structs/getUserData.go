package main

import (
	"errors"
	"fmt"
	"strconv"
)

func getUserData(userData string) (string, error) {
	var dummy string
	fmt.Println(userData)
	fmt.Scanln(&dummy)
	if dummy == "" {
		return "nill", errors.New("please enter the input dont let the input empty")
	} else {
		return dummy, nil

	}
}
func getUserPhone(str string) (int64, error) {
	fmt.Println(str)
	var num int64

	fmt.Scanln(&num)
	numText := strconv.FormatInt(num, 10)
	n := len(numText)
	if n < 10 {
		return 0, errors.New("please enter the mobile number of 10 digits")
	}
	return num, nil

}
