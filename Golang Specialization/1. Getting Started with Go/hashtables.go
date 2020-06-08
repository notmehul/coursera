package main

import "fmt"

func main() {
    //HASH TABLES BOIZZ (MAPS)
	// Hash table is basically a dictionary lol
	var idMap map[string]int

	idMap = make(map[string]int)

	idMap["notmehul"] = 69

	// one way to have maps

	idMap2 := map[string]int{
		"mehul":           1,
		"Mehul":           2,
		"notmehul":        3,
		"not_mehul":       4,
		"totallyNotMehul": 5,
		"mehulmao":        6
	}

	//fmt.Println(idMap2["Mehul"])
	//fmt.Println(idMap["notmehul"])

	x, y := idMap2["Mehul"]
	fmt.Println(x, y) //y here is true because the key exists in the hash

	for key, val := range idMap2 {
		fmt.Println(key, val)
	}
}