package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Printf("Enter a string: ")
	fmt.Scanln(&str)
	str = strings.ToLower(str)

	if str[0] == 'i' && str[len(str)-1] == 'n' && strings.Contains(str, "a") {
		fmt.Printf("Found\n")
	} else {
		fmt.Printf("Not Found\n")
	}
}
