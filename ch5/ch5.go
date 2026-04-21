package main

import (
	"fmt"
)

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// )

// func add(i int, j int) (int, error) { return i + j, nil }

// func sub(i int, j int) (int, error) { return i - j, nil }

// func mul(i int, j int) (int, error) { return i * j, nil }

// func div(i int, j int) (int, error) {
// 	if j == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return i / j, nil
// }

// var opMap = map[string]func(int, int) (int, error){
// 	"+": add,
// 	"-": sub,
// 	"*": mul,
// 	"/": div,
// }

// func main() {
// 	expressions := [][]string{
// 		{"2", "+", "3"},
// 		{"2", "-", "3"},
// 		{"2", "*", "3"},
// 		{"2", "/", "3"},
// 		{"2", "%", "3"},
// 		{"two", "+", "three"},
// 		{"5"},
// 		{"2", "/", "0"},
// 	}
// 	for _, expression := range expressions {
// 		if len(expression) != 3 {
// 			fmt.Println("invalid expression:", expression)
// 			continue
// 		}
// 		p1, err := strconv.Atoi(expression[0])
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		op := expression[1]
// 		opFunc, ok := opMap[op]
// 		if !ok {
// 			fmt.Println("unsupported operator:", op)
// 			continue
// 		}
// 		p2, err := strconv.Atoi(expression[2])
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		result, error := opFunc(p1, p2)
// 		if error != nil {
// 			fmt.Println(error)
// 			continue
// 		}
// 		fmt.Println(result)

// 	}
// }

// func fileLen(name string) (int, error) {
// 	file, err := os.Open(name)

// 	if err != nil {
// 		return 0, err
// 	}

// 	defer file.Close()

// 	data := make([]byte, 2048)
// 	allCount := 0
// 	for {
// 		count, err := file.Read(data)

// 		allCount += count

// 		if err != nil {
// 			if err != io.EOF {
// 				break
// 			}
// 			return 0, err
// 		}
// 	}

// 	return allCount, nil
// }

// func main() {
// 	count, err := fileLen("./Makefile")

// 	fmt.Println(count)
// 	fmt.Println(err)

// }

func prefixer(input string) func(string) string {
	return func(word string) string {
		return input + " " + word
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
