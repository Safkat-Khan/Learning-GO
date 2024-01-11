package main

import (
	"fmt"
)

func main() {
	var s string
	var i int

	fmt.Print("Enter Your Name: ")
	fmt.Scan(&s)
	fmt.Print("Enter Your Age: ")
	fmt.Scanln(&i)
	fmt.Println("Your Name is ", s, "Your Age is", i)
}
