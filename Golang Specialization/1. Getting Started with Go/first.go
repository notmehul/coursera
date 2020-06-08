package main

import "fmt"

//HOW STRUCT IS DONE LOL

type person struct {
	name  string
	age   int
	phone string
}

func main() {

	//POINTERS

	var x int = 1
	var y int
	var ip *int //pointer

	ip = &x
	y = *ip
	fmt.Print(y)
	fmt.Print(ip)

}
