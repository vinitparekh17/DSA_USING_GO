package main

import (
	"fmt"
)

func main() {
	ar := []int{7, 2, 1, 6, 8, 5, 3, 4}
	QuickSort(ar, 0, len(ar)-1)
	fmt.Println(ar)
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		pivot := Partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func Partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			Swapping(arr, i, j)
		}
	}
	Swapping(arr, i+1, high)
	return i + 1
}

func Swapping(arr []int, i int, j int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}
