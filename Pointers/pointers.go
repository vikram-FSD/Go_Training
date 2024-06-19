package main

import "fmt"

func main() {
	age := 32 //regular variable
	fmt.Println("Age: ", age)
	agePointer := &age
	adultAge(agePointer)
	fmt.Println(age)
}
func adultAge(age *int) {
	// return *age - 18
	*age = *age - 18

}
