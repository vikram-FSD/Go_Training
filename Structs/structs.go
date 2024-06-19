package main

func main() {
	// 	Name, err1 := getUserData("Please Enter YourName: ")
	// 	if err1 != nil {
	// 		fmt.Println("Name Error: ", err1)
	// 		return
	// 	}
	// 	Pass, err2 := getUserData("Please Enter Password: ")
	// 	if err2 != nil {
	// 		fmt.Println("Password Error: ", err2)
	// 		return
	// 	}
	// 	Dob, err3 := getUserData("Please Enter Your BithDate (DD/MM/YYYY): ")
	// 	if err3 != nil {
	// 		fmt.Println("Date of birth Error: ", err3)
	// 		return
	// 	}
	// 	var data user.UserData

	// 	fmt.Println(data)

	// 	data = user.UserData{
	// 		UserName:  Name,
	// 		Password:  Pass,
	// 		BirthDate: Dob,
	// 		CreatedAt: time.Now(),
	// 	}
	// 	outputUserData(data)
	// 	// printemp()

	// }

	// func outputUserData(u user.UserData) {
	// 	fmt.Println("\n", "UserName: ", u.UserName, "\n", "Birthdate: ", u.BirthDate, "\n", "Password: ", u.Password, "\n", "CreatedAt: ", u.CreatedAt)
	// createNewStruct()
	var name customType = "vikram"
	name.log()
}
