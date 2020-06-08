package main

import "fmt"

func main() {
	//STRUCTSS
	p1 := new(person) //initializes everything to zero value

	fmt.Println(p1.age)
	// 0

	p2 := person{name: "Mehul", age: 18, phone: "9717945717"}

	fmt.Println(p2.name)
}
