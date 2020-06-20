package main

import (
	"fmt"

	demo "github.com/nurali-techie/module-demo"
)

func main() {
	userName := demo.UserName("Nurali", "Virani")
	fmt.Printf("userName=%s\n", userName)
}
