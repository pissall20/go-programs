package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal struct {
//	food, locomotion, noise string
//}

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	// food, locomotion, noise string
}

type Bird struct {
	// food, locomotion, noise string
}

type Snake struct {
	// food, locomotion, noise string
}

func (a *Cow) Eat() {
	fmt.Println("grass")
}

func (a *Cow) Move() {
	fmt.Println("walk")
}

func (a *Cow) Speak() {
	fmt.Println("moo")
}

func (a *Bird) Eat() {
	fmt.Println("worms")
}

func (a *Bird) Move() {
	fmt.Println("fly")
}

func (a *Bird) Speak() {
	fmt.Println("peep")
}

func (a *Snake) Eat() {
	fmt.Println("mice")
}

func (a *Snake) Move() {
	fmt.Println("slither")
}

func (a *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	dict := make(map[string]AnimalInterface)
	for {
		// var command, animal, action string
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		split_input := strings.Split(strings.TrimSpace(input), " ")
		command, animal, action := split_input[0], split_input[1], split_input[2]
		switch command {
		case "newanimal":
			switch action {
			case "cow":
				dict[animal] = new(Cow)
			case "bird":
				dict[animal] = new(Bird)
			case "snake":
				dict[animal] = new(Snake)
			}
			fmt.Println("Created it!")
		case "query":
			obj, ok := dict[animal]
			if ok {
				switch action {
				case "eat":
					obj.Eat()
				case "move":
					obj.Move()
				case "speak":
					obj.Speak()
				}
			} else {
				fmt.Println("Not found. Please add the animal with 'newanimal'")
			}
		}
	}
}
