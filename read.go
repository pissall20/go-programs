package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

const maxLength = 20

type Name struct {
	fname, lname string
}

func (n *Name) Set(firstName, lastName string) {
	var rs []rune

	n.fname = firstName
	if len(firstName) > maxLength {
		rs = []rune(firstName)
		n.fname = string(rs[:maxLength])
	}

	n.lname = lastName
	if len(lastName) > maxLength {
		rs = []rune(lastName)
		n.lname = string(rs[:maxLength])
	}
}

func main(){
	nameSlice := make([]Name, 0)

	var fileName string
	// Get the file name and it's contents
	fmt.Println("Enter the path of the file with Name contents:")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), " ")
			name1 := Name{}
			name1.Set(s[0], s[1])
			nameSlice = append(nameSlice, name1)
		}
	}
	for _, name := range nameSlice {
		fmt.Println(name.fname, name.lname)
	}
}
