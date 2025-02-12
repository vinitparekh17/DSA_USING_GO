## Fibonacci Series using Dynamic Programming

> The reason for considering the Fibonacci series is that it is the simplest problem to understand the concept of dynamic programming.

### Problem Statement

Given a number `n`, print the `n`-th Fibonacci number.

### Solution

`main.go` contains the solution to the problem with four different approaches:

1. **Recursive Approach (Classic Approach)**: A straightforward recursive solution that has exponential time complexity.
2. **DP Approach (Top-Down Approach)**: Uses memoization to store previously computed values, reducing the time complexity to linear.
3. **DP Approach (Bottom-Up Approach)**: Builds the solution from the base cases up to the desired `n`, also achieving linear time complexity.
4. **DP Approach (Bottom-Up Approach with Space Optimization)**: Further optimizes the bottom-up approach by using only constant space.