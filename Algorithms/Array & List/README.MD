# Arrays

Arrays are a data structure that stores a collection of elements of the same type. Arrays are indexed, which means that each element can be accessed by its index. The index of the first element is 0, the index of the second element is 1, and so on.

To declare an array in Go:

```
arr := [5]int{}
```

where:

* `arr` is the name of the array
* `5` is the size of the array
* `int` is the type of the elements in the array

```go
fmt.Println(arr[0])
```

## Graphical Representation of an Array:

| Index | 0 | 1 | 2 | 3 | 4 |
|---|---|---|---|---|---|
| Value | 0 | 0 | 0 | 0 | 0 |

## Methods of Arrays:

Arrays in Go have a number of methods that can be used to manipulate the array. These methods include:

* `Len()`: Returns the length of the array.
* `Cap()`: Returns the capacity of the array.
* `At(index)`: Returns the element at the specified index.
* `Get(index)`: Returns a copy of the element at the specified index. 
* `Set(index, value)`: Sets the element at the specified index to the specified value.
* `Append(value)`: Appends the specified value to the end of the array.
* `Copy(dst, src)`: Copies the elements from the specified source array to the destination array.
* `Swap(a, b)`: Swaps the values of the specified elements.
* `Sort()`: Sorts the elements of the array.
* `Reverse()`: Reverses the order of the elements of the array.

# Slices

Slices are a data structure that stores a collection of elements of the same type. Slices are indexed, which means that each element can be accessed by its index. The index of the first element is 0, the index of the second element is 1, and so on.

## Deference between Arrays and Slices:

| Feature | Array | Slice |
|---|---|---|
| Size | Fixed | Dynamic |
| Initialization | `array_name := [size]element_type` | `slice_name := make([]element_type, size)` |
| Appending elements | Not possible | Possible |
| Resizing | Not possible | Possible |
| Memory usage | More | Less |
| Performance | Faster | Slower |
| Use cases | When We need a fixed-size data structure | When We need a dynamic-size data structure |

# Linked Lists

Linked lists are a data structure that stores a collection of elements, where each element is linked to the next element. Linked lists are not indexed, which means that We cannot access elements by their index.

To declare a linked list in Go, We use the following syntax:

```go
list_name := list.New()
```

## Graphical Representation of a Linked List:

| Value | Next |
|---|---|
| 1 | 2 |
| 2 | 3 |
| 3 | 4 |


## Methods of Linked Lists:

Linked lists in Go have a number of methods that can be used to manipulate the linked list. These methods include:

* `Len()`: Returns the length of the linked list.
* `Front()`: Returns the first element of the linked list.
* `Back()`: Returns the last element of the linked list.
* `PushFront(value)`: Adds the value to the front of the linked list.
* `PushBack(value)`: Adds the value to the back of the linked list.
* `InsertBefore(value)`: Adds the value to the linked list before the specified element.
* `InsertAfter(value)`: Adds the value to the linked list after the specified element.
* `Remove(e)`: Removes the specified element from the linked list.
* `MoveToFront(e)`: Moves the specified element to the front of the linked list.
* `MoveToBack(e)`: Moves the specified element to the back of the linked list.

We can iterate over a linked list using a for loop:

```go
for e := list_name.Front(); e != nil; e = e.Next() {
	fmt.Println(e.Value)
}
```