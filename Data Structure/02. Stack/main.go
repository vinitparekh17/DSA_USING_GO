package main

import "fmt"

type Stack struct {
	items []int
}

var size int

func (s *Stack) Push() {
	var i int
	if IsFull(s) {
		fmt.Println("Stack is full")
	} else {
		fmt.Println("Enter a number to push to the stack: ")
		fmt.Scanln(&i)
		s.items = append(s.items, i)
		fmt.Println("Pushed", i, "to the stack: ", *s)
	}
}

func (s *Stack) Show() {
	fmt.Println(*s)
}

func (s *Stack) Pop() {
	if IsEmpty(s) {
		fmt.Println("Stack is empty")
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		fmt.Println("Popped", item, "from the stack: ", *s)
	}
}

func (s *Stack) Peek() {
	if IsEmpty(s) {
		fmt.Println("Stack is empty")
	}
	fmt.Println(s.items[len(s.items)-1])
}

func IsFull(s *Stack) bool {
	return len(s.items) == size
}

func IsEmpty(s *Stack) bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{items: []int{}}
	fmt.Println("Declare the size of the stack: ")
	fmt.Scanln(&size)
	for {
		var input int
		fmt.Println("Enter a number to choose an option: ")
		fmt.Println("1. Push\n2. Pop\n3. Peek\n4. Show\n5. Exit")
		fmt.Scanln(&input)
		switch input {
		case 1:
			stack.Push()
		case 2:
			stack.Pop()
		case 3:
			stack.Peek()
		case 4:
			stack.Show()
		case 5:
			return
		}
	}
}
