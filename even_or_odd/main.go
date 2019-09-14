package main

import "fmt"

// func main() {
// 	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	for index, number := range numbers {
// 		if number%2 == 0 {
// 			fmt.Println(index, "even")
// 		} else {
// 			fmt.Println(index, "odd")
// 		}
// 	}
// }

func main() {
	var s [11]int

	for i, _ := range s {
		s[i] = i
		if s[i]%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("oddd")
		}
	}
}
