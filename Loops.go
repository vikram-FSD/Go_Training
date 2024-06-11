package main

import "fmt"

func printValue(i int, text string) {
	loop(i, text)
}
func loop(i int, text string) {
	for i := i; i <= 3; i++ {
		fmt.Println(i, ".", text)
	}
}
