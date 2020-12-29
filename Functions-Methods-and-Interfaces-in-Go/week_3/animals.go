package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
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
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "X" {
			break
		}

		input = strings.ToLower(input)
		result := strings.Split(input, " ")
		animal := result[0]
		action := result[1]

		a, exist := zoo[animal]
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
			fmt.Printf(action)
		}

	}
}
