package main

import "fmt"

func main() {
	fmt.Println("Assignment Operators in Go")

	// Basic assignment
	var x int = 40
	fmt.Println("Initial value of x:", x)

	// Addition assignment
	x += 5 // x = x + 5
	fmt.Println("x += 5:", x)

	// Subtraction assignment
	x = 40 // Reset value
	x -= 5 // x = x - 5
	fmt.Println("x -= 5:", x)

	// Multiplication assignment
	x = 40 // Reset value
	x *= 5 // x = x * 5
	fmt.Println("x *= 5:", x)

	// Division assignment
	x = 40 // Reset value
	x /= 5 // x = x / 5
	fmt.Println("x /= 5:", x)

	// Modulus assignment
	x = 40 // Reset value
	x %= 5 // x = x % 5
	fmt.Println("x %= 5:", x)
}
