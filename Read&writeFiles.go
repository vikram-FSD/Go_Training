package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalanceTOFile(text string, bal float64) {
	balanceText := fmt.Sprintf("%s%.1f", text, bal)
	os.WriteFile("balance.txt", []byte(balanceText), 0655)
}
func getBalanceFromFile() (float64, error) {
	file := "accBal.txt"
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, errors.New("failed to find balance file")
	}
	str := string(data)
	bal, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New("failed to parse the stored balance file")
	}
	return bal, nil

}
