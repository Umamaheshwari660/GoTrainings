// Go data types are broadly categorized as:
//
// 1. Basic Types
//    - bool       : true or false
//    - string     : sequence of characters (UTF-8 encoded)
//    - numeric    : integers, floating-point numbers, complex numbers
//
// 2. Aggregate Types
//    - array      : fixed-length collection of elements of the same type
//    - struct     : collection of fields with different types
//
// 3. Reference Types
//    - pointer    : holds the memory address of a value
//    - slice      : dynamic view over an array
//    - map        : key-value data structure
//    - function   : callable function type
//    - channel    : used for communication between goroutines
//
// 4. Interface Types
//    - interface  : defines behavior through method sets
//
// This file demonstrates examples of Basic Types.

package main

import "fmt"

func main() {
	// Boolean type
	// Stores true or false values
	var isReady bool = true

	// String type
	// Immutable sequence of bytes (UTF-8 encoded)
	var message string = "Hello, Go!"

	// Integer type
	// 'int' size depends on the architecture (32-bit or 64-bit)
	var num int = 42

	// Floating-point type
	// float64 is the default floating-point type
	var price float64 = 19.99

	// Complex number type
	// complex128 consists of two float64 values: real and imaginary
	var comp complex128 = complex(2, 3)

	// Output values
	fmt.Printf("bool: %v\n", isReady)
	fmt.Printf("string: %v\n", message)
	fmt.Printf("int: %v\n", num)
	fmt.Printf("float64: %v\n", price)
	fmt.Printf("complex128: %v\n", comp)
}
