package main

import "fmt"

func mapping() {
	websites := map[string]string{

		"Google":    "https://google.com",
		"aws":       "https://aws.com",
		"instagram": "https://instagram.com",
	}
	fmt.Println(websites)
	delete(websites, "Google")
	fmt.Println(websites)
}
