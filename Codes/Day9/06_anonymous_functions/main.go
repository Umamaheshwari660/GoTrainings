package main

import "fmt"

// ============================================================
// ANONYMOUS FUNCTIONS (FUNCTIONS WITHOUT A NAME)
// ============================================================

func main() {
	fmt.Println("=== Anonymous Functions ===")

	// 1. Assigning an anonymous function to a variable
	greet := func(name string) {
		fmt.Println("Hello,", name)
	}
	greet("Go lang")
	greet("Gopher")

	// 2. Immediately Invoked Function Execution (IIFE)
	//    Defining and calling a function in one step.
	result := func(a, b int) int {
		return a + b
	}(3, 7) // ← the parentheses () immediately invoke the function
	fmt.Println("IIFE result (3+7) =", result)

	// 3. Anonymous function without storing result
	func() {
		fmt.Println("This runs immediately (IIFE with no return)")
	}()

	// 4. Using anonymous functions as closures
	//    A closure is a function that captures variables from its surrounding scope.
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Closure counter:")
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// 5. Passing anonymous function as an argument
	numbers := []int{1, 2, 3, 4, 5}
	// Example: using an anonymous function with a custom "map" operation
	transformed := transform(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squares:", transformed)

	// 6. Returning an anonymous function from another function
	multiplyBy := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := multiplyBy(2)
	triple := multiplyBy(3)
	fmt.Println("double(5) =", double(5)) // 10
	fmt.Println("triple(5) =", triple(5)) // 15
}

// transform applies a function to each element of a slice.
func transform(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}
