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
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func RandomizedPartition(arr []int, low int, high int) int {
	randomIndex := rand.Intn(high-low+1) + low
	Swap(arr, randomIndex, high)
	return Partition(arr, low, high)
}

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
