package main

import "fmt"

func closure() {
	num := increment()
	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	anotherNum := increment()
	fmt.Println(anotherNum())
}
func increment() func() int {
	num := 0

	return func() int {
		num = num + 1
		return num
	}
}
