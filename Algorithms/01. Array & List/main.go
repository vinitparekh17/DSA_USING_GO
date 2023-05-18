package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	arr := [5]int{3, 1, 2, 5, 4}
	slice := arr[:]  // slice of the arr
	fmt.Println(arr) // simple arr

	// arr operations
	slice = append(slice, 6) // append 6 to the arr
	fmt.Println(slice)

	// arr = append(arr[:2], arr[3:]...) // delete 3rd element from the arr
	fmt.Println(slice)

	// sorting the arr
	sort.Ints(slice)
	fmt.Println(slice)

	// get
	fmt.Println(slice[0]) // get the first element of the arr

	// set
	slice[0] = 0 // set the first element of the arr to 0
	fmt.Println(slice)

	// reverse
	slice = reverse(slice)
	fmt.Println(slice)

	// initilizing a linked list
	linkedList := list.New()
	fmt.Println(linkedList)

	// adding elements to the linked list
	linkedList.PushFront(1)
	fmt.Println(linkedList.Front().Value)        // prints the first element
	linkedList.PushBack(2)                       // adding second elm.
	fmt.Println(linkedList.Back().Value)         // prints the last element
	fmt.Println(linkedList.Front().Next().Value) // Next() returns the next element in the list

	//Note: address of the next element will be shown if the element is not present or I don't use Value after Next()

	// Let's push another list into the first list
	anotherList := list.New()
	anotherList.PushBack(3)
	anotherList.PushBack(4)

	linkedList.PushBackList(anotherList)
	fmt.Println("After pushing another list into the first list")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// delete middle elem. from list
	linkedList.Remove(linkedList.Front().Next()) // removes 2
	fmt.Println("After removing 2")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// insert elem. in mid
	linkedList.InsertAfter(2, linkedList.Front()) // add 2 at position 2 ezy :)
	fmt.Println("Insert after first elem..")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func reverse(arr []int) []int {
	reversed := make([]int, len(arr))
	for i := 0; i < len(arr)/2; i++ {
		reversed[i], reversed[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return reversed
}
