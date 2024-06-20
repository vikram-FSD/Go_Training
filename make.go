package main

import "fmt"

func makeFunction() {
	type floatMap map[string]float64
	// userNames := make([]string, 2)
	// userNames = append(userNames, "Virkam")
	// userNames = append(userNames, "Vijay")
	// userNames = append(userNames, "Baalu")
	// userNames = append(userNames, "Berlin")
	// userNames = append(userNames, "John")
	// fmt.Println(userNames)
	movies := make(floatMap, 3)
	movies["Leo"] = 5.0
	movies["Darbar"] = 3.4
	movies["Jersey"] = 10.0
	movies["Ravening"] = 5
	movies["Dangal"] = 10.0
	fmt.Println(movies)

}
