package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{5, 7, 4, 6, 3, 2, 1}
	sorted, err := SelectonSort(arr)
	HandleError(err)
	for _, e := range sorted {
		fmt.Println(e)
	}
}

func SmallestIndex(arr []int) (int, error) {
	small := arr[0]
	index := 0

	if len(arr) == 0 {
		return 0, errors.New("empty error not allowed, jackass")
	}

	for i := range arr {
		if arr[i] < small {
			small = arr[i]
			index = i
		}
	}

	return index, nil
}

func SelectonSort(arr []int) ([]int, error) {
	newArr := []int{}
	// logic
	return newArr, nil
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
