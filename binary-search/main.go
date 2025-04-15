package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	array := processDictionary("words_alpha.txt")
	word := "zoo"

	index, guesses := binarySearch(array, word)
	fmt.Printf("index of %s: %d\titerations using binary search: %d\n", word, index, guesses)

	index, guesses = sequentialSearch(array, word)
	fmt.Printf("index of %s: %d\titerations using sequential search: %d\n", word, index, guesses)
}

func processDictionary(filename string) (dictionary []string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	for line := range strings.SplitSeq(string(bytes), "\n") {
		before, word, found := strings.Cut(line, "\t")
		if found {
			dictionary = append(dictionary, word)
		} else {
			dictionary = append(dictionary, strings.TrimSpace(before))
		}
	}

	return dictionary
}

func sequentialSearch(array []string, item string) (index int, guesses int) {
	for index, str := range array {
		guesses += 1
		if str == item {
			return index, guesses
		}
	}

	return -1, guesses
}

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
