package main

import "fmt"

// ============================================================
// RECURSION: A FUNCTION THAT CALLS ITSELF
// ============================================================

// Recursive function to print numbers from n to 5
func printNumbers(n int) {
	if n > 5 { // base condition: stop when n exceeds 5
		return
	}
	fmt.Println("printNumbers:", n)
	printNumbers(n + 1) // recursive call
}

// Recursive function to calculate sum of numbers from 1 to num
func sumRecursive(num int) int {
	// base condition
	if num <= 0 {
		return 0
	}
	// recursive case: num + sum of (num-1)
	return num + sumRecursive(num-1)
}

// Recursive factorial function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Iterative equivalent of sumRecursive (for comparison)
func sumIterative(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println("=== Recursion Examples ===")

	// Print numbers recursively
	fmt.Println("Printing numbers from 1 to 5 using recursion:")
	printNumbers(1)

	// Sum using recursion
	fmt.Println("\nSum of 1 to 10 (recursive):", sumRecursive(10))

	// Sum using iteration
	fmt.Println("Sum of 1 to 10 (iterative):", sumIterative(10))

	// Factorial
	fmt.Println("Factorial of 5 (recursive):", factorial(5))

	// Base condition explanation:
	// Every recursive function must have a base condition that stops the recursion.
	// Without it, the function would call itself indefinitely, causing a stack overflow.

	// Call stack visualization for sumRecursive(3):
	// sumRecursive(3) → 3 + sumRecursive(2)
	//                   → 2 + sumRecursive(1)
	//                         → 1 + sumRecursive(0)
	//                               → 0 (base)
	//                         ← 1 + 0 = 1
	//                   ← 2 + 1 = 3
	//             ← 3 + 3 = 6
}
