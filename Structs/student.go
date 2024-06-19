package main

import "fmt"

type StudentData struct {
	StudName     string
	StudRollNo   int64
	StudEmail    string
	StudMobile   int64
	StudLocation string
	Laptopspec
}
type Laptopspec struct {
	LapName string
	Ram     string
}

func new(name, email, location string, rollNo, mobile int64) StudentData {
	return StudentData{
		StudName:     name,
		StudRollNo:   rollNo,
		StudEmail:    email,
		StudMobile:   mobile,
		StudLocation: location,
		Laptopspec: Laptopspec{
			LapName: "HP",
			Ram:     "16GB",
		},
	}
}
func createNewStruct() {

	name, err := getUserData("StudentName: ")
	if err != nil {
		fmt.Println(err)
	}
	email, err := getUserData("Email: ")
	if err != nil {
		fmt.Println(err)
	}
	loc, err := getUserData("Student_Location: ")
	if err != nil {
		fmt.Println(err)
	}
	roll, err := getUserPhone("Student_rollNo: ")
	if err != nil {
		fmt.Println(err)
	}
	mob, err := getUserPhone("Student_mobile: ")
	if err != nil {
		fmt.Println(err)
	}

	data := new(name, email, loc, roll, mob)
	fmt.Println(data)
}
