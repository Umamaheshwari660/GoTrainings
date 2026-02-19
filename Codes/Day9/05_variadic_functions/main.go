package main

import "fmt"

// ============================================================
// VARIADIC FUNCTIONS
// ============================================================

// A variadic function accepts a variable number of arguments.
// The parameter is declared with an ellipsis (...) before the type.
// Inside the function, it behaves as a slice of that type.

// -------------------- EXAMPLE 1: SUM OF INTEGERS --------------------
// NumberSum adds any number of int values.
func NumberSum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// -------------------- EXAMPLE 2: FIXED + VARIADIC --------------------
// greetWithNumbers takes a greeting string followed by any number of ints.
func greetWithNumbers(greeting string, nums ...int) {
	fmt.Print(greeting)
	for _, n := range nums {
		fmt.Print(" ", n)
	}
	fmt.Println()
}

// -------------------- EXAMPLE 3: CONCATENATE STRINGS --------------------
// concat joins strings with a separator.
func concat(separator string, parts ...string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += separator
		}
		result += part
	}
	return result
}

// -------------------- EXAMPLE 4: FIND MAXIMUM --------------------
// max returns the largest integer from the provided numbers.
// If no numbers are given, it prints a warning and returns 0.
func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for _, val := range nums[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// -------------------- EXAMPLE 5: PASS A SLICE TO VARIADIC --------------------
// demonstrateSlicePassing shows how to pass an existing slice using "..."
func demonstrateSlicePassing() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", numbers)
	// To pass a slice, append ... after it
	sum := NumberSum(numbers...)
	fmt.Println("Sum of slice (passed with ...):", sum)
}

// -------------------- EXAMPLE 6: EMPTY VARIADIC CALL --------------------
// demonstrateEmptyCall shows calling a variadic function with no arguments.
func demonstrateEmptyCall() {
	fmt.Println("Calling NumberSum() with no arguments:")
	fmt.Println("Result:", NumberSum()) // Output: 0
}

// ============================================================
// MAIN – CLEAN AND ORGANIZED
// ============================================================
func main() {
	fmt.Println("🔹 VARIADIC FUNCTIONS IN GO 🔹")

	// Example 1: Basic sum
	fmt.Println("--- 1. Basic Sum ---")
	fmt.Println("NumberSum(2, 6)      =", NumberSum(2, 6))    // 8
	fmt.Println("NumberSum(2, 6, 4)   =", NumberSum(2, 6, 4)) // 12
	fmt.Println("NumberSum()           =", NumberSum())       // 0
	fmt.Println()

	// Example 2: Fixed + variadic
	fmt.Println("--- 2. Fixed String + Variadic Ints ---")
	greetWithNumbers("Numbers:", 1, 2, 3)
	greetWithNumbers("Only one:", 42)
	greetWithNumbers("No numbers:") // empty variadic part
	fmt.Println()

	// Example 3: Concatenate strings
	fmt.Println("--- 3. String Concatenation ---")
	fmt.Println("concat(', ', 'apple', 'banana', 'cherry') =",
		concat(", ", "apple", "banana", "cherry"))
	fmt.Println("concat(' - ', 'one', 'two') =",
		concat(" - ", "one", "two"))
	fmt.Println("concat(' ... ') with no strings =",
		concat(" ... ")) // empty result
	fmt.Println()

	// Example 4: Maximum value
	fmt.Println("--- 4. Maximum of Variadic Ints ---")
	fmt.Println("max(4, 7, 2, 9, 3) =", max(4, 7, 2, 9, 3)) // 9
	fmt.Println("max(100)           =", max(100))           // 100
	fmt.Println("max()              =", max())              // warning + 0
	fmt.Println()

	// Example 5: Passing a slice
	fmt.Println("--- 5. Passing a Slice to a Variadic Function ---")
	demonstrateSlicePassing()
	fmt.Println()

	// Example 6: Empty variadic call
	fmt.Println("--- 6. Empty Variadic Call ---")
	demonstrateEmptyCall()
	fmt.Println()

	/*
		📌 KEY POINTS:
			• Variadic parameter is a slice inside the function
			• It must be the last parameter (if mixed with fixed ones)
			• Use slice... to pass an existing slice.
			• Functions cannot be overloaded in Go; variadic functions provide flexibility
	*/
}
