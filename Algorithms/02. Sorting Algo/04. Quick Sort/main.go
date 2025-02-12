package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := []int{7, 2, 1, 6, 8, 5, 3, 4}
	QuickSort(ar, 0, len(ar)-1)
	fmt.Println(ar)
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		pivot := RandomizedPartition(arr, low, high)
		// The first recursive call sorts the subarray that contains the elements that are smaller than the pivot.
		QuickSort(arr, low, pivot-1)
		// The second recursive call sorts the subarray that contains the elements that are greater than the pivot.
		QuickSort(arr, pivot+1, high)
	}
}

// using random pivot to avoid worst case scenario of O(n^2)
// because if the pivot is the smallest or largest element in the array
// then the partition function will have to run n times
// which will result in O(n^2)
func RandomizedPartition(arr []int, low int, high int) int {
	// this rand.Intn is using pseudo-random number generator (PRNG) algorithm under the hood
	// check out https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/math/rand/rand.go;l=358
	// for more info :)
	randomIndex := rand.Intn(high-low+1) + low
	Swap(arr, randomIndex, high)
	return Partition(arr, low, high)
}

// partitioning the array into two halves and returning the index of the pivot
func Partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			Swap(arr, i, j)
		}
	}
	Swap(arr, i+1, high)
	return i + 1
}

func Swap(arr []int, i int, j int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}
