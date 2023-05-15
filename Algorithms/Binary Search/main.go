package main

import (
	"fmt"
	"time"
)

func main() {
	array := make([]int, 50000)
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
	fmt.Println("Enter a number between 1 and 50,000 to search for:")
	var input int
	fmt.Scanln(&input)
	BinarySearch(array, input)
}

func BinarySearch(arr []int, input int) {
	start := time.Now()
	fmt.Println("Searching for", input, "...")
	attempts := 0
	low := arr[0]
	high := len(arr)
	mid := (low + high) / 2
	for low <= high {
		attempts++
		if arr[mid] == input {
			fmt.Println("Found", input, "at index", mid)
			break
		} else if arr[mid] < input {
			low = mid + 1
		} else if arr[mid] > input {
			high = mid - 1
		}
		mid = (low + high) / 2
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms\nAttempts:", attempts)
}
