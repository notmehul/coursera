package main

import "fmt"

func main() {
	// LOOPS

	for i := 0; i < 10; i++ {
		fmt.Printf("hi")
	}
	// prints hi 10 times like a fucking creep lmao
	y := [6]int{1, 2, 3, 4, 5, 6}

	for i, v := range y {
		fmt.Printf("%d %d,", i, v)
	}
}
