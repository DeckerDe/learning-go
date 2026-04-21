package main

import "fmt"

const value = 14

func main() {
	// Exercise 1
	i := 20
	
	f := float64(i)

	fmt.Println(i)
	fmt.Println(f)

	// Exercise 2

	var (
		k int     = value
		j float64 = value
	)

	fmt.Println(j)
	fmt.Println(k)

	// Exercise 3

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
