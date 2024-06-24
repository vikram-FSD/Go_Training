package main

import "fmt"

func anonymous() {

	func() {
		fmt.Println("Iam anonymous!")
	}()
}
