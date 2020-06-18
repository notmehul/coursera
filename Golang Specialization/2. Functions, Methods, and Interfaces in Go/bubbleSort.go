package main

import "fmt"

//BubbleSort sorts the array
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			swap(arr, i)
		}
	}
}

func swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main() {
	arr := make([]int, 0, 10)
	fmt.Println(arr)
	fmt.Println("enter the array to be sorted")

	for i := 0; i < 10; i++ { //adds the input to the array
		var inp int
		fmt.Scan(&inp)
		arr = append(arr, inp)
	}
	BubbleSort(arr)
	fmt.Println(arr)
}
