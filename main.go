package main

import (
	"fmt"

	demo "github.com/nurali-techie/module-demo/v2"
)

func main() {
	userName := demo.UserName("Nurali", "Virani", "Barkatali")
	fmt.Printf("userName=%s\n", userName)

	password := demo.GeneratePassword()
	fmt.Printf("password=%s\n", password)

	email := demo.GenerateEmail(userName)
	fmt.Printf("email=%s\n", email)
}
