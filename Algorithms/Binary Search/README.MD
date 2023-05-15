# Binary search

## Senario

Suppose you are searching for word from dictionary. You can start from the first page and search for the word. If you don't find it, you can move to the next page and search again. You can repeat this process until you find the word. This is called **linear search**.

But there is a better way to do this. You can open the dictionary from the middle and check if the word is in the first half or the second half. If the word is in the first half, you can open the first half of the dictionary and repeat the process. If the word is in the second half, you can open the second half of the dictionary and repeat the process. This is called **binary search**.

Note that the dictionary should be sorted in order to use binary search.

## Implementation

[Binary search](main.go) is implemented in Go.

1) What is the time complexity of binary search?
ans: O(log n)

2) What is the space complexity of binary search?
ans: O(1)