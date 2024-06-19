package main

import (
	"fmt"

	"example.com/basics-practice/fileOperations"
)

func callwriteTOfilefun() {
	var str string
	fmt.Println("enter the content to be write to the file: ")
	fmt.Scan(&str)
	fileOperations.WriteTotheFile(str)
}
