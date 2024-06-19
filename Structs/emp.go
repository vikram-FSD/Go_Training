package main

import (
	"fmt"
	"time"
)

// Creating the structure for employee
type emp struct {
	ename       string
	eId         string
	empPhone    int64
	city        string
	currentTime time.Time
}

func newEmployee(Name string, Id string, phone int64, loc string, ct time.Time) emp {
	return emp{ename: Name, eId: Id, empPhone: phone, city: loc, currentTime: ct}
}

func printemp() {
	Name, err := getUserData("Enter Emp Name: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	Id, err := getUserData("Enter Emp ID: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	loc, err := getUserData("Enter City: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	phone, err := getUserPhone("Enter Phone: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ct := time.Now()

	var data = newEmployee(Name, Id, phone, loc, ct)
	printOutput(data)

}
func printOutput(v emp) {
	v = emp{ename: v.ename, eId: v.eId, city: v.city, currentTime: v.currentTime, empPhone: v.empPhone}
	fmt.Println("\n", v.eId, "\n", v.ename, "\n", v.city, "\n", v.currentTime, "\n", v.empPhone)
}
