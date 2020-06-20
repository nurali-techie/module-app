package main

import (
	"fmt"

	demo "github.com/nurali-techie/module-demo"
	demoV2 "github.com/nurali-techie/module-demo/v2"
)

func main() {
	userName := demo.UserName("Nurali", "Virani")
	fmt.Printf("userName=%s\n", userName)

	password := demo.GeneratePassword()
	fmt.Printf("password=%s\n", password)

	email := demoV2.GenerateEmail(userName)
	fmt.Printf("email=%s\n", email)
}
