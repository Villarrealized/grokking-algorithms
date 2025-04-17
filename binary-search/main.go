package main

import (
	"fmt"
	"grokking-algorithms/files"
)

func main() {

	// fmt.Println(sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	array := files.ProcessDictionary("../words_alpha.txt")
	word := "zoo"

	index, guesses := binarySearch(array, word)
	fmt.Printf("index of %s: %d\titerations using binary search: %d\n", word, index, guesses)

	index, guesses = sequentialSearch(array, word)
	fmt.Printf("index of %s: %d\titerations using sequential search: %d\n", word, index, guesses)
}

// O(n)
func sequentialSearch(array []string, item string) (index int, guesses int) {
	for index, str := range array {
		guesses += 1
		if str == item {
			return index, guesses
		}
	}

	return -1, guesses
}

// O(logn)
func binarySearch(array []string, item string) (index int, guesses int) {
	low := 0
	high := len(array) - 1
	guesses = 0

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		guesses += 1

		if guess == item {
			return mid, guesses
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, guesses
}

func sum(ints []int) (total int) {
	if len(ints) == 0 {
		return 0
	}

	return ints[0] + sum(ints[1:])
}
