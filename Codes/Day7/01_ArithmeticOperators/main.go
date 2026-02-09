package main

import "fmt"

func main() {
	fmt.Println("Arithmetic Operators in Go")

	// Basic arithmetic operators
	a := 100
	b := 300

	fmt.Println("a =", a, "b =", b)

	// Addition
	sum := a + b
	fmt.Printf("Addition: %d + %d = %d\n", a, b, sum)

	// Subtraction
	difference := a - b
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, difference)

	// Multiplication
	product := a * b
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, product)

	// Division (integer division for integers)
	quotient1 := a / b
	fmt.Printf("Integer Division: %d / %d = %d\n", a, b, quotient1)

	quotient2 := b / a
	fmt.Printf("Integer Division: %d / %d = %d\n", b, a, quotient2)

	// Modulus (remainder)
	remainder1 := a % b
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, remainder1)

	remainder2 := b % a
	fmt.Printf("Modulus: %d %% %d = %d\n", b, a, remainder2)

	// Floating point arithmetic
	fmt.Println("\n--- Floating Point Arithmetic ---")

	aFloat := 100.0
	bFloat := 300.0

	fmt.Println("Floating point division:", aFloat/bFloat) // 0.333333...
	fmt.Println("Floating point division:", bFloat/aFloat) // 3.0

	// Mixed type arithmetic (needs explicit conversion)
	fmt.Println("\n--- Mixed Type Arithmetic ---")
	var intValue int = 50
	var floatValue float64 = 25.5

	// This won't compile: fmt.Println(intValue + floatValue)
	// Need explicit conversion
	result1 := float64(intValue) + floatValue
	fmt.Printf("int + float = %.2f\n", result1)

	result2 := intValue + int(floatValue) // Truncates float
	fmt.Printf("int + int(float) = %d\n", result2)

	// Increment and Decrement operators
	fmt.Println("\n--- Increment and Decrement ---")

	x := 20
	fmt.Println("Original x =", x)

	// Post-increment
	x++ // x = x + 1
	fmt.Println("After x++ =", x)

	// Post-decrement
	x-- // x = x - 1
	fmt.Println("After x-- =", x)

	// Important: In Go, ++ and -- are statements, not expressions
	// This is NOT allowed: y := x++
	// You must do:
	y := x
	x++
	fmt.Println("y =", y, "x =", x)

	// Pre-increment and pre-decrement don't exist in Go
	// ++x and --x are NOT valid in Go

	// Compound assignment with arithmetic
	fmt.Println("\n--- Compound Assignment Operators ---")

	value := 10
	fmt.Println("Initial value =", value)

	value += 5 // value = value + 5
	fmt.Println("After value += 5 =", value)

	value -= 3 // value = value - 3
	fmt.Println("After value -= 3 =", value)

	value *= 2 // value = value * 2
	fmt.Println("After value *= 2 =", value)

	value /= 4 // value = value / 4
	fmt.Println("After value /= 4 =", value)

	value %= 3 // value = value % 3
	fmt.Println("After value %%= 3 =", value)

	// Working with different integer types
	fmt.Println("\n--- Different Integer Types ---")

	var int8Val int8 = 100
	var int16Val int16 = 300
	var int32Val int32 = 500

	// Need explicit conversion for arithmetic between different types
	sum1 := int16(int8Val) + int16Val
	fmt.Printf("int8 + int16 = %d\n", sum1)

	sum2 := int32(int8Val) + int32Val
	fmt.Printf("int8 + int32 = %d\n", sum2)

	// Division by zero
	fmt.Println("\n--- Division by Zero ---")

	numerator := 10
	// denominator := 0

	// This would cause a runtime panic:
	// result := numerator / denominator

	// Always check for zero before division
	denominator := 0
	if denominator != 0 {
		result := numerator / denominator
		fmt.Println("Division result:", result)
	} else {
		fmt.Println("Cannot divide by zero!")
	}

	// Modulus with negative numbers
	fmt.Println("\n--- Modulus with Negative Numbers ---")

	fmt.Printf("10 %% 3 = %d\n", 10%3)
	fmt.Printf("-10 %% 3 = %d\n", -10%3)
	fmt.Printf("10 %% -3 = %d\n", 10%-3)
	fmt.Printf("-10 %% -3 = %d\n", -10%-3)

	// The sign of the result follows the dividend (first number)

	// Order of operations
	fmt.Println("\n--- Order of Operations ---")

	result := 2 + 3*4 // Multiplication before addition
	fmt.Println("2 + 3 * 4 =", result)

	result = (2 + 3) * 4 // Parentheses change order
	fmt.Println("(2 + 3) * 4 =", result)

	// Complex expression
	result = 10 + 2*3 - 4/2 + 5%3
	// Step by step: 2*3=6, 4/2=2, 5%3=2
	// Then: 10 + 6 - 2 + 2 = 16
	fmt.Println("10 + 2*3 - 4/2 + 5%3 =", result)
}
