// Write a program which prompts the user to first enter a name,
//  and then enter an address. Your program should create a map and
// add the name and address to the map using the keys “name” and “address”,
//  respectively. Your program should use Marshal() to create a JSON object from the map,
//   and then your program should print the JSON object.

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	name    string
// 	address string
// }

// func main() {
// 	var person Person
// 	m := make(map[string]string)

// 	fmt.Printf("\nEnter a name: ")
// 	fmt.Scanln(&person.name)

// 	fmt.Printf("\nEnter a address: ")
// 	fmt.Scanln(&person.address)

// 	m["name"] = person.name
// 	m["address"] = person.address

// 	barr, _ := json.Marshal(m)

// 	fmt.Println(string(barr))
// }

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var inpt string
	output := make(map[string]string)
	fmt.Printf("please enter a name:")
	fmt.Scan(&inpt)
	output["name"] = inpt
	fmt.Printf("please enter a address:")
	fmt.Scan(&inpt)
	output["address"] = inpt
	jsonSt, _ := json.Marshal(output)
	fmt.Println(string(jsonSt))
}
