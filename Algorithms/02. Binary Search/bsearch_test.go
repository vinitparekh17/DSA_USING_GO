package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	slice := make([]int, 50000)
	input := 6353
	resultIndex := BinarySearchTest(slice, input)
	if resultIndex == -1 {
		fmt.Println("Input is greater than the last element in the array")
	} else if resultIndex == -2 {
		fmt.Println("Input is less than the first element in the array")
	} else if resultIndex != 6352 {
		t.Error("BinarySearch() failed")
	}
}

func BinarySearchTest(arr []int, input int) int {
	if input > arr[len(arr)-1] {
		return -1
	} else if input < arr[0] {
		return -2
	}
	low := 0
	high := len(arr)
	mid := (high + low) / 2
	result := 0
	for low <= high {
		if arr[mid] < input {
			low = mid + 1
		} else if arr[mid] > input {
			low = mid - 1
		} else {
			result = mid
		}
	}
	return result
} // result: PASS ^_^

func BenchmarkBinarySearch(b *testing.B) {
	slice := make([]int, 50000)
	input := 6353
	for i := 0; i < b.N; i++ {
		BinarySearchTest(slice, input)
	}
} // result: Approx 0.500s below for 50000 elements :)
