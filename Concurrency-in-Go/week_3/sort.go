package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	part := 4
	maxSize := 50
	numbers := make([]int, 0, maxSize)

	fmt.Println("Enter integers separeted with space (less than 50)")
	inputReader := bufio.NewReader(os.Stdin)
	strInput, _ := inputReader.ReadString('\n')
	splitIntegers := strings.Fields(strInput)
	ok := checkInputSize(len(splitIntegers), part, maxSize)

	if ok {
		for _, value := range splitIntegers {
			value, err := strconv.Atoi(value)
			if err == nil {
				numbers = append(numbers, value)
			}
		}
		arrayPartitions := make([][]int, 4)
		partition(numbers, part, arrayPartitions)

		var wg sync.WaitGroup
		wg.Add(part)
		for i := 0; i < part; i++ {
			go sortPartition(&arrayPartitions[i], &wg)
		}
		wg.Wait()

		res := arrayPartitions[0]
		for _, sortedPartition := range arrayPartitions[1:part] {
			res = merge(res, sortedPartition)
		}

		fmt.Println(res)
	}

}

func checkInputSize(size, part, maxSize int) bool {
	if size < part {
		fmt.Printf("There less than %d integers , ending", part)
		return false
	}

	if size > maxSize {
		fmt.Printf("There more than %d integers , ending", maxSize)
		return false
	}

	return true
}

func convertToInt(strInput []string, numbers []int) {
	for _, value := range strInput {
		num, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Printf(" %s is not an integer , ending", value)
			return
		}
	}
}

func partition(numbers []int, part int, arrayPartitions [][]int) [][]int {

	size := len(numbers) / part
	arrayPartitions[0] = numbers[:size]
	arrayPartitions[1] = numbers[size : size*2]
	arrayPartitions[2] = numbers[size*2 : size*3]
	arrayPartitions[3] = numbers[size*3 : len(numbers)]
	return arrayPartitions
}

func sortPartition(partitionIntegers *[]int, wg *sync.WaitGroup) {
	fmt.Println("Unsorted partition: ", *partitionIntegers)
	sort.Ints(*partitionIntegers)
	wg.Done()
}

func merge(leftArray, rightArray []int) []int {
	mergedSize := len(leftArray) + len(rightArray)
	mergedArray := make([]int, mergedSize, mergedSize)
	leftIndex, rightIndex := 0, 0

	for i := 0; i < mergedSize; i++ {
		if leftIndex >= len(leftArray) {
			mergedArray[i] = rightArray[rightIndex]
			rightIndex++
		} else if rightIndex >= len(rightArray) {
			mergedArray[i] = leftArray[leftIndex]
			leftIndex++
		} else if leftArray[leftIndex] < rightArray[rightIndex] {
			mergedArray[i] = leftArray[leftIndex]
			leftIndex++
		} else {
			mergedArray[i] = rightArray[rightIndex]
			rightIndex++
		}
	}
	return mergedArray
}
