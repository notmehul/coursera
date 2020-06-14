package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := []int{}
	var x string
	for true {
		fmt.Println("Hello there, enter an integer(press X to exit):")
		fmt.Scan(&x)
		if x == "X" {
			break
		} else {
			ap, err := strconv.Atoi(x)
			arr = append(arr, ap)
			if err != nil {
				fmt.Println("Wrong input")
				continue
			}
		}
		sort.Ints(arr[:])
		fmt.Println(arr)
	}
}
