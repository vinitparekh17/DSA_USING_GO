## Concept of selection sort

<b>"Selection sort is a sorting algorithm that selects the smallest element from an unsorted list in each iteration and places that element at the beginning of the unsorted list."</b>

<hr/>
<i>Can't understood? Let's understand by assuming a scenario.</i>

Suppose, you are playing cards with your friends. You have a deck of cards and you want to arrange them in ascending order. What will you do?

Simple way is to find the smallest card from the deck and put it aside. Then find the second smallest card and put it aside. Repeat this process until all the cards are put aside. This is called selection sort.

Now let's put on our developer's hat and see the logic behind the above scenario and try to relate it with the selection sort algorithm.

So running simple search on the cards will take O(n) time. And we have to do this for n cards. So the total time complexity will be O(n<sup>2</sup>).

Because outer searching (loop) will be executed n times to find the smallest card and inner searching (loop) will be executed n times to find the smallest card. So the total time complexity will be O(n<sup>2</sup>).

## Use cases of selection sort

- [x] Names in a phone book
- [x] Train numbers by arrival time
- [x] Emails (newest to oldest)

Well, after getting the idea of selection sort, you might have a question that as you go through the operation, the number of unsorted elements decreases and ends up with one element at the end so how the time complexity is O(n<sup>2</sup>)?

<b>The answer has to do with constants in Big O notation!</b>

Here is the quicl gist of it:
You check n elements then n - 1  n - 2 … 2, 1. On average, you check a list that has 1/2 × n elements. The runtime is O(n × 1/2 × n). But constants like 1/2 are ignored in Big O notation. but why? Because Big O notation only cares about the general trend of the runtime as the input grows. So the runtime is O(n<sup>2</sup>).

## Selection sort implementation

[Check it out](main.go)
```go
func SelectonSort(arr []int) ([]int, error) {
	newArr := []int{}
	for range arr {
		smallest, err := SmallestIndex(arr) // finds the smallest elem. from remaining elems.
		HandleError(err)
		newArr = append(newArr, arr[smallest]) // adds smallest elem.
		arr = append(arr[:smallest], arr[smallest+1:]...) // removes smallest elem.
	}
	return newArr, nil
}
```

In above code as we discussed earlier that we will have to use nested loops in this algorithm, inner to find the smallest element and outer to add that element to the new array and remove it from the old array.



## Selection sort complexity
O(n<sup>2</sup>)

