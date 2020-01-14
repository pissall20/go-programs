package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	data := make(map[string]string)
	var name, address string 
	fmt.Println("Please enter the name:")
	fmt.Scan(&name)

	fmt.Println("Please enter the address:")
	fmt.Scan(&address)
	data["name"] = name
	data["address"] = address
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Some problem occurred.\n")
	}
	fmt.Println(data)
	fmt.Println(string(jsonStr))

}
