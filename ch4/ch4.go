package main

// func main() {
// 	var numbers []int

// 	for i := 0; i <= 100; i++ {
// 		number := rand.IntN(100)

// 		numbers = append(numbers, number)
// 	}

// 	for k := range numbers {
// 		switch {
// 		case k%2 == 0 && k%3 == 0:
// 			println("Six")
// 		case k%2 == 0:
// 			println("Two")
// 		case k%3 == 0:
// 			println("Three")
// 		default:
// 			println("Never Mind")
// 		}

// 	}
// }

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		println(total)
	}
}
