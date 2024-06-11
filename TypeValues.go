package main

import "fmt"

func typevalues() {
	num := 100.12345
	fmt.Printf("%v,%T\n", num, num)
	value := 2000.00
	resultValue := fmt.Sprintf("Printed_Value: %.0f\n", value)
	fmt.Print(resultValue)

}
