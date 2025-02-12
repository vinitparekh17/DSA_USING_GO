package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{5, 7, 4, 6, 3, 1, 2}
	sorted, err := SelectonSort(arr)
	HandleError(err)
	fmt.Println(sorted)
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
	for range arr {
		smallest, err := SmallestIndex(arr)
		HandleError(err)
		// add the smallest element to the new array
		newArr = append(newArr, arr[smallest])
		// remove the smallest element from the old array to run SmallestIndex() recursively and don't get the same smallest element
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr, nil
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
