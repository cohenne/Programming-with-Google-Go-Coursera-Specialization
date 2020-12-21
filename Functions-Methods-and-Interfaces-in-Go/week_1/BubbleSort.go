package main

import (
	"fmt"
)

func main() {
	length := 10
	fmt.Printf("Please enter %d numbers: \n", length)
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanln(&numbers[i])
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}

func BubbleSort(numbers []int) {
	for i, _ := range numbers {
		for j := 0; j < len(numbers)-1-i; j++ {
			// fmt.Println(j)
			if numbers[j] > numbers[j+1] {
				swap(numbers, j)
			}
		}
	}
}

func swap(arr []int, index int) {
	arr[index], arr[1+index] = arr[index+1], arr[index]
}
