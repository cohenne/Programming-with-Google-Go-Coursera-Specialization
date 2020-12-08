// Your program will define a name struct which has two fields,
//  fname for the first name, and lname for the last name.
//  Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
//  Your program will successively read each line of the text file and
//  create a struct which contains the first and last names found in the file.
//  Each struct created will be added to a slice, and after all lines have been read from the file,
//  your program will have a slice containing one struct for each line in the file.
//  After reading all lines from the file, your program should iterate through your
//  slice of structs and print the first and last names found in each struct.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var names = make([]Name, 0)

	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var name = new(Name)

		line := scanner.Text()
		line = strings.TrimSuffix(line, "\n")

		name.fname = strings.Fields(line)[0]
		name.lname = strings.Fields(line)[1]

		names = append(names, *name)
	}

	for _, n := range names {
		fmt.Printf("fname: %s, lname: %s\n", n.fname, n.lname)
	}
}
