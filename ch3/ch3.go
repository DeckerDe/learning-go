package main

import "fmt"

func main() {
	// Exercise 1
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	x := greetings[:2]
	y := greetings[1:5]
	z := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Exercise 2
	message := "Hi 👧 and 👦"

	runes := []rune(message)

	fmt.Println(string(runes[3]))

	// Exercise 3

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	first := Employee{
		"John",
		"Doe",
		1,
	}

	second := Employee{
		firstName: "Bob",
		lastName:  "Builder",
		id:        2,
	}

	var third Employee

	third.firstName = "Jessica"
	third.lastName = "Meyers"
	third.id = 3

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
