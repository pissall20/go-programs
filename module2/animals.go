package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	dict := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	for {
		var animal, action string
		fmt.Print(">")
		fmt.Scan(&animal)
		fmt.Print(">")
		fmt.Scan(&action)
		switch action {
		case "eat":
			dict[animal].Eat()
		case "move":
			dict[animal].Move()
		case "speak":
			dict[animal].Speak()
		}
	}
}
