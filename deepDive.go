package main

import "fmt"

type Transform func(int) int

func mains() {
	numbers := []int{1, 4, 6}
	// tripleNumbers := []int{1, 2, 3}
	// transformNumbers(numbers, double)
	// transformNumbers(tripleNumbers, triple)
	dummy := getTranformFunction(numbers)
	fmt.Println(dummy)
}

// func transformNumbers(numbers []int, transform Transform) {
// 	dNum := []int{}
// 	for _, value := range numbers {
// 		dNum = append(dNum, transform(value))
// 	}
// 	fmt.Println(dNum)

// }
func getTranformFunction(num []int) Transform {
	if num[0] == 2 {
		return double
	} else {
		return triple

	}

}
func double(numbers int) int {
	return numbers * 2
}
func triple(numbers int) int {
	return numbers * 3
}
