package main

import (
	"fmt"
	"time"
)

func main() {
	n := 10
	fmt.Printf("Solving Fibonacci problem using Dynamic Programming to understand the Complexities\n\n")

	// Classic recursive approach
	classic = 0
	start := time.Now()
	result1 := fibonacci(n)
	duration1 := time.Since(start)
	fmt.Printf("Classic Recursive: %d\n", result1)
	fmt.Printf("Function calls: %d\n", classic)
	fmt.Printf("Time taken: %v\n\n", duration1)

	// Top-down DP (Memoization)
	i1 = 0
	start = time.Now()
	result2 := fib1(n, make(map[int]int))
	duration2 := time.Since(start)
	fmt.Printf("Top-Down DP (Memoization): %d\n", result2)
	fmt.Printf("Function calls: %d\n", i1)
	fmt.Printf("Time taken: %v\n\n", duration2)

	// Bottom-up DP (Tabulation)
	i2 = 0
	start = time.Now()
	result3 := fib2(n)
	duration3 := time.Since(start)
	fmt.Printf("Bottom-Up DP (Tabulation): %d\n", result3)
	fmt.Printf("Iterations: %d\n", i2)
	fmt.Printf("Time taken: %v\n\n", duration3)

	i3 = 0
	start = time.Now()
	result4 := fib3(n)
	duration4 := time.Since(start)
	fmt.Printf("Bottom-Up DP (Tabulation) Constant Space: %d\n", result4)
	fmt.Printf("Iterations: %d\n", i3)
	fmt.Printf("Time taken: %v\n", duration4)

}

// fibonacci implements the classic recursive approach to calculate Fibonacci numbers.
// Time Complexity: O(2^n) - exponential, as it recalculates same values multiple times
// Space Complexity: O(n) - due to recursion stack depth
// Advantages: Simple to understand and implement
// Disadvantages: Very inefficient for larger numbers due to repeated calculations
var classic int

func fibonacci(n int) int {
	classic++ // func Counter
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// fib1 implements Top-Down Dynamic Programming (Memoization) approach.
// Time Complexity: O(n) - each value calculated only once
// Space Complexity: O(n) - for storing memo map plus recursion stack
// Advantages: Avoids recalculation of values, more efficient than classic approach
// Disadvantages: Still uses recursion which can cause stack overflow for very large n
var i1 int

func fib1(n int, memo map[int]int) int {
	i1++ // func Counter
	if n <= 1 {
		return n
	}
	if val, found := memo[n]; found {
		return val
	}
	memo[n] = fib1(n-1, memo) + fib1(n-2, memo)
	return memo[n]
}

// fib2 implements Bottom-Up Dynamic Programming (Tabulation) approach.
// Time Complexity: O(n) - linear
// Space Complexity: O(n) - for storing dp array
// Advantages: Most efficient, no recursion overhead, predictable space usage
// Disadvantages: Uses more memory than necessary (can be optimized to O(1) space)
var i2 int

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		i2++ // iterations Counter
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// fib3 implements same as above but optimized to O(1) space, since only last 2 value should be remembered
// Time Complexity: O(n) - linear
// Space Complexity: O(1)
// Advantages: Best version, no recursion, constant space usage
// Disadvantages: N/A
var i3 int

func fib3(n int) int {
	if n <= 1 {
		return n
	}
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		i3++ // Iterations Counter
		current := prev1 + prev2
		prev2, prev1 = prev1, current
	}
	return prev1
}
