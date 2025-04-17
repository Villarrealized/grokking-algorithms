package main

import (
	"fmt"
	"grokking-algorithms/files"
	"math/rand"
)

func main() {
	unsorted := files.ProcessDictionary("../common-words.txt")
	sorted, operations := quickSort(unsorted)
	fmt.Printf("sorted %d elements in %d operations using quick sort\n", len(sorted), operations)

	sorted, operations = insertionSort(unsorted)
	fmt.Printf("sorted %d elements in %d operations using insertion sort\n", len(sorted), operations)

	sorted, operations = selectionSort(unsorted)
	fmt.Printf("sorted %d elements in %d operations using selection sort\n", len(sorted), operations)

	sorted, operations = bubbleSort(unsorted)
	fmt.Printf("sorted %d elements in %d operations using bubble sort\n", len(sorted), operations)
}

// O(n^2)
// uses a two pointer approach where the left pointer represents the current insertion point
// then the left pointer value is compared to every other value
// if some other value is smaller, it is swapped
// after the first pass the smallest value will be at index 0 and the left pointer advances to the next item in the list
func selectionSort(array []string) (sorted []string, operations int) {
	length := len(array)
	sorted = make([]string, length)
	copy(sorted, array)

	for left := range length {
		min := left
		for right := left + 1; right < length; right++ {
			if sorted[right] < sorted[min] {
				min = right
			}
			operations += 1
		}
		// swap
		sorted[min], sorted[left] = sorted[left], sorted[min]
	}
	return sorted, operations
}

// O(n^2)
// better perf than selection sort, one of the best n^2 algos and better than quicksort on very small data sets
// starts with the 2nd element as all arrays of length 1 are sorted by definition
// the right pointer searches backwards and inserts the current element in sorted order
// each iteration all elements to the left of the left pointer are in sorted order
func insertionSort(arr []string) (sorted []string, operations int) {
	length := len(arr)
	sorted = make([]string, length)
	copy(sorted, arr)

	for left := 1; left < length; left++ {
		for right := left; right > 0 && sorted[right-1] > sorted[right]; right-- {
			sorted[right], sorted[right-1] = sorted[right-1], sorted[right]
			operations += 1
		}
	}
	return sorted, operations
}

// O(n^2)
// worst of the n^2 sorting algos unless the array is already sorted or mostly sorted
// bubble up each max value to the end of the array by comparing pairs i, i + 1
// each time reduce the array length by 1 until no values have been swapped
// after first pass the largest value will be at the end of the array
func bubbleSort(array []string) (sorted []string, operations int) {
	length := len(array)
	sorted = make([]string, length)
	copy(sorted, array)

	swapped := false

	for ok := true; ok; ok = swapped {
		swapped = false
		for i := range length - 1 {
			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
				swapped = true
			}
			operations += 1
		}
		length -= 1
	}
	return sorted, operations
}

// average case O(n*logn), worst case O(n^2)
func quickSort(array []string) (sorted []string, operations int) {

	var sort func(array []string) []string
	sort = func(array []string) []string {
		length := len(array)
		if length <= 1 {
			return array
		}

		// Choose a random pivot index
		pivotIndex := rand.Intn(length)

		// Swap the pivot with the first element
		array[0], array[pivotIndex] = array[pivotIndex], array[0]

		pivot := array[0]
		less := make([]string, 0, length/2)
		greater := make([]string, 0, length/2)

		for i := 1; i < length; i++ {
			s := array[i]
			if s <= pivot {
				less = append(less, s)
			} else {
				greater = append(greater, s)
			}
			operations += 1
		}
		return append(append(sort(less), pivot), sort(greater)...)
	}
	return sort(array), operations
}
