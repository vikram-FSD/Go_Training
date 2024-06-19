package main

import "fmt"

type customType string //Creating other custom types and Adding methods

func (text customType) log() {
	fmt.Println(text)
}
