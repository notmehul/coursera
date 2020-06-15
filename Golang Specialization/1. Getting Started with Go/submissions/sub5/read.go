package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type name struct {
	fname []byte
	lname []byte
}

func main() {

	var file string

	fmt.Print("Enter the filename:")
	fmt.Scan(&file)

	data, _ := ioutil.ReadFile(file)
	content := string(data)

	words := strings.Fields(content)

	names := make([]name, 0)

	for i := 0; i < len(words)-1; i = i + 2 {
		n := name{fname: []byte(words[i]), lname: []byte(words[i+1])}

		names = append(names, n)
	}
	for _, n := range names {
		fmt.Printf("first name: %s, last name: %s\n", n.fname, n.lname)
	}
}
