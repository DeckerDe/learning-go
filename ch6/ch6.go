package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	var persons = make([]Person, 1000000)
	//var persons []Person
	for i := 0; i < 1000000; i++ {
		person := Person{
			FirstName: "aoba",
			LastName:  "kk",
			Age:       16,
		}
		persons = append(persons, person)
	}

	fmt.Println(len(persons))

}
