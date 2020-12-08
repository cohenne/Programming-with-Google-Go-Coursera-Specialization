package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)

	for {
		var input string
		// var cnt int
		fmt.Println("Enter an integer to be added to an array or X for exit")
		fmt.Scanf("%s", &input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			continue
		}

		// if cnt < 3 {
		// 	sli[cnt] = num
		// } else {
		// 	sli = append(sli, num)
		// }

		sli = append(sli, num)

		sort.Ints(sli)
		fmt.Println(sli)

	}
}
