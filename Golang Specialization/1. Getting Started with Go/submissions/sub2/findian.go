package main

import (
	"fmt"
	"strings"
)

func main() {
	var toFind = "a"
	var toCheck string
	fmt.Scan(&toCheck)
	TorF := strings.Contains(toCheck, toFind)
	if toCheck[0] == 'i' && toCheck[len(toCheck)-1] == 'n' && TorF {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found")
	}
}
