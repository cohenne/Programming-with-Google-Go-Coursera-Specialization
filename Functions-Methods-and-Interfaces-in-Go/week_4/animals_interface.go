// package main

// import (
// 	"fmt"
// )

// type Animal struct {
// 	food       string
// 	locomotion string
// 	noise      string
// }

// type AnimalIterface interface {
// 	Eat()
// 	Move()
// 	Speak()
// }

// func (a Animal) Eat() {
// 	fmt.Println(a.food)
// }

// func (a Animal) Move() {
// 	fmt.Println(a.locomotion)
// }

// func (a Animal) Speak() {
// 	fmt.Println(a.noise)
// }

// func main() {
// 	zoo := make(map[string]Animal)

// 	zoo["cow"] = Animal{"grass", "walk", "moo"}
// 	zoo["bird"] = Animal{"worms", "fly", "peep"}
// 	zoo["snake"] = Animal{"mice ", "slither", "hsss"}

// 	for {
// 		var command, animal, action string
// 		fmt.Print(">")
// 		fmt.Scan(&command, &animal, &action)

// 		if command == "X" {
// 			break
// 		}

// 		if command == "query" {

// 			a, exist := zoo[animal]
// 			if !exist {
// 				fmt.Printf("Error! animal %s is not in my list", a)
// 			}
// 			switch action {
// 			case "eat":
// 				a.Eat()
// 			case "move":
// 				a.Move()
// 			case "speak":
// 				a.Speak()
// 			default:
// 				fmt.Printf("Error")
// 			}
// 		} else if command == "newanimal" {
// 			zoo[animal] = zoo[action]
// 			fmt.Println("Created it!")
// 		} else {
// 			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
// 		}
// 	}
// }
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal props
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow specs
type Cow struct {
	food       string
	locomotion string
	noise      string
}

//Bird specs
type Bird struct {
	food       string
	locomotion string
	noise      string
}

//Snake specs
type Snake struct {
	food       string
	locomotion string
	noise      string
}

//NewAnimal any
type NewAnimal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter three strings space seperated, for example:\n>newanimal myanimal cow\nOR\n>query cow eat\n")
	fmt.Print(">")
	for scanner.Scan() {
		usrReq := scanner.Text()
		info := strings.Split(usrReq, " ")
		switch strings.ToLower(info[0]) {
		case "newanimal":
			var a NewAnimal
			_, err := a.addAnimal(info[2])
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf(">Created it!\n")
			}
		case "query":
			switch strings.ToLower(info[1]) {
			case "cow":
				c := Cow{"grass", "walk", "moo"}
				switch info[2] {
				case "eat":
					c.Eat()
				case "move":
					c.Move()
				case "speak":
					c.Speak()
				default:
					fmt.Println(">no data found, please enter any of eat, move or speak")
				}
			case "bird":
				b := Bird{"worms", "fly", "peep"}
				switch strings.ToLower(info[2]) {
				case "eat":
					b.Eat()
				case "move":
					b.Move()
				case "speak":
					b.Speak()
				default:
					fmt.Println(">no data found, please enter any of eat, move or speak")
				}
			case "snake":
				s := Snake{"mice", "slither", "hsss"}
				switch info[2] {
				case "eat":
					s.Eat()
				case "move":
					s.Move()
				case "speak":
					s.Speak()
				default:
					fmt.Println(">no data found, please enter any of eat, move or speak")
				}
			default:
				fmt.Println(">no data found, please enter any of cow, bird or snake")
			}
		default:
			fmt.Printf(">%s: Command not found\n", info[0])
		}
		fmt.Print(">")
	}
}

func (a NewAnimal) addAnimal(animalType string) (NewAnimal, error) {
	switch animalType {
	case "cow":
		a.food = "grass"
		a.locomotion = "walk"
		a.noise = "moo"
	case "bird":
		a.food = "worms"
		a.locomotion = "fly"
		a.noise = "peep"
	case "snake":
		a.food = "mice"
		a.locomotion = "slither"
		a.noise = "hsss"
	default:
		return a, fmt.Errorf(">no data found, please enter any of cow, bird or snake")
	}
	return a, nil
}

//Eat food
func (c Cow) Eat() {
	fmt.Printf(">%s\n", c.food)
}

//Move Locomotion
func (c Cow) Move() {
	fmt.Printf(">%s\n", c.locomotion)
}

//Speak noise
func (c Cow) Speak() {
	fmt.Printf(">%s\n", c.noise)
}

//Eat food
func (b Bird) Eat() {
	fmt.Printf(">%s\n", b.food)
}

//Move Locomotion
func (b Bird) Move() {
	fmt.Printf(">%s\n", b.locomotion)
}

//Speak noise
func (b Bird) Speak() {
	fmt.Printf(">%s\n", b.noise)
}

//Eat food
func (s Snake) Eat() {
	fmt.Printf(">%s\n", s.food)
}

//Move Locomotion
func (s Snake) Move() {
	fmt.Printf(">%s\n", s.locomotion)
}

//Speak noise
func (s Snake) Speak() {
	fmt.Printf(">%s\n", s.noise)
}
