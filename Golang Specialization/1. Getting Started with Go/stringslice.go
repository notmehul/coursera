package main

import "fmt"

func main() {
	//STRING SLICES
	arr := [...]string{"a", "b", "c", "d", "e", "f"}
	s1 := arr[1:5]
	fmt.Println(len(s1), cap(s1))
	// if we take the length of the slice, we will get 4'
	// if we take the capacity of the slice, we will get capacity
	//that's because 5 is the max value it can hold as the size of the arr is 5 lol
	sli := make([]string, 10) //make function is used to make the slice lol
	fmt.Println(sli)          // outputs empty string
	sli = append(sli, "lololol waste of memory")
	fmt.Println(sli)
}
