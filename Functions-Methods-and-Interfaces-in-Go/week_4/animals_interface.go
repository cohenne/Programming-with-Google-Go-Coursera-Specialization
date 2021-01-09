package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

type AnimalIterface interface {
	Eat()
	Move()
	Speak()
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
	zoo := make(map[string]Animal)

	zoo["cow"] = Animal{"grass", "walk", "moo"}
	zoo["bird"] = Animal{"worms", "fly", "peep"}
	zoo["snake"] = Animal{"mice ", "slither", "hsss"}

	for {
		var command, animal, action string
		fmt.Print(">")
		fmt.Scan(&command, &animal, &action)

		if command == "X" {
			break
		}

		if command == "query" {

			a, exist := zoo[ani]
			if !exist {
				fmt.Printf("Error! animal %s is not in my list", a)
			}
			switch action {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Printf("Error")
			}
		} else if command == "newanimal" {
			zoo[animal] = zoo[action]
			fmt.Println("Created it!")
		} else {
			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
		}
	}
}
