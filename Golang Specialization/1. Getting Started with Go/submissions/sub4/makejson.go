package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)
	var nam string
	var addr string

	fmt.Println("enter the name and address")
	fmt.Scan(&nam, &addr)

	person[nam] = addr

	arr, _ := json.Marshal(person)

	fmt.Println("Encoded data : ")
	fmt.Println(arr)
	fmt.Println("Decoded data : ")
	fmt.Println(string(arr))
}
