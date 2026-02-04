package main

import "fmt"

/*
ZERO VALUES IN GO

In Go, every variable is initialized with a "zero value" if no explicit value is provided.
This is different from some languages where uninitialized variables contain garbage/undefined values.

Zero values ensure:
1. Memory safety - no undefined behavior from uninitialized memory
2. Predictable behavior - programs behave consistently
3. Simpler code - less boilerplate initialization

Each type has its own zero value:
- Numeric types: 0
- Boolean: false
- Strings: "" (empty string)
- Pointers, slices, maps, channels, functions, interfaces: nil
- Arrays/structs: each element/field gets its respective zero value
*/

func main() {
	fmt.Println("=== ZERO VALUES IN GO ===")
	fmt.Println("Every type in Go has a default 'zero value' when declared without initialization.")
	fmt.Println("This eliminates undefined behavior and makes code more predictable.")

	// ========== BASIC TYPES ==========
	fmt.Println("--- BASIC TYPES ---")

	// Boolean: zero value is false
	var b bool // false
	fmt.Printf("bool:      %#v\t// false - logical 'off' state\n", b)

	// String: zero value is empty string
	var s string // ""
	fmt.Printf("string:    %#v\t// \"\" - empty string, valid for all string operations\n", s)

	// Integer: zero value is 0
	var i int // 0
	fmt.Printf("int:       %#v\t// 0 - neutral element for addition\n", i)

	// Float: zero value is 0.0
	var f float64 // 0.0
	fmt.Printf("float64:   %#v\t// 0.0 - represents no quantity\n", f)

	// Complex: zero value is (0+0i)
	var c complex128 // (0+0i)
	fmt.Printf("complex:   %#v\t// (0+0i) - zero in complex plane\n", c)

	// ========== REFERENCE TYPES ==========
	fmt.Println("\n--- REFERENCE TYPES (all nil by default) ---")
	fmt.Println("Reference types are 'nil' when declared, meaning they don't point to any data.")
	fmt.Println("Attempting to use nil references without initialization causes runtime panic.")

	// Pointer: zero value is nil (doesn't point to anything)
	var p *int // nil
	fmt.Printf("pointer:   %#v\t// nil - safe to check, dangerous to dereference\n", p)

	// Slice: zero value is nil (not an empty slice!)
	var sl []int // nil
	fmt.Printf("slice:     %#v\t// nil - different from empty slice ([]int{})\n", sl)

	// Map: zero value is nil (cannot be used until initialized with make())
	var m map[string]int // nil
	fmt.Printf("map:       %#v\t// nil - must initialize with make() before use\n", m)

	// Channel: zero value is nil (cannot send/receive)
	var ch chan int // nil
	fmt.Printf("channel:   %#v\t// nil - closed channel, cannot send or receive\n", ch)

	// Function: zero value is nil (cannot be called)
	var fn func() // nil
	fmt.Printf("function:  %#v\t// nil - cannot be called until assigned a function\n", fn)

	// ========== AGGREGATE TYPES ==========
	fmt.Println("\n--- AGGREGATE TYPES ---")
	fmt.Println("Aggregate types initialize each element/field to its respective zero value.")

	// Array: each element gets zero value of the element type
	var arr [3]int // [0, 0, 0]
	fmt.Printf("array:     %#v\t// All 3 elements initialized to 0\n", arr)

	// Struct: each field gets its respective zero value
	var st struct { // All fields zero
		x    int
		y    float64
		name string
	}
	fmt.Printf("struct:    %#v\t// Each field gets its type's zero value\n", st)

	// ========== PRACTICAL IMPLICATIONS ==========
	fmt.Println("\n=== PRACTICAL IMPLICATIONS ===")

	// 1. SAFE DEFAULT BEHAVIOR
	fmt.Println("1. Safe Default Behavior:")
	fmt.Println("   Zero values provide safe defaults:")
	var count int
	var total float64
	var found bool

	// These are safe to use immediately
	average := total / float64(count) // 0/0 = NaN? No, Go returns 0 for 0/0 with floats
	fmt.Printf("   Average of %v items with total %v: %v\n", count, total, average)
	fmt.Printf("   Search result found: %v (false means 'not found' by default)\n", found)

	// 2. INITIALIZATION PATTERNS
	fmt.Println("\n2. Common Initialization Patterns:")

	// Maps must be initialized before use
	var counts map[string]int
	if counts == nil {
		fmt.Println("   ✓ Map is nil - this is expected and safe to check")
		fmt.Println("   Initializing map with make()...")
		counts = make(map[string]int)
	}
	counts["apples"] = 5 // Now safe to use
	fmt.Printf("   Map after initialization: %v\n", counts)

	// Arrays are immediately usable
	var scores [5]int // Already initialized to [0,0,0,0,0]
	scores[0] = 100   // Safe immediately
	fmt.Printf("   Array ready immediately: %v\n", scores)

	// 3. ZERO VALUE VS EXPLICIT EMPTY VALUE
	fmt.Println("\n3. Zero Value vs Explicit Empty Value:")

	// Nil slice vs empty slice
	var nilSlice []int             // nil slice
	emptySlice := []int{}          // empty slice (non-nil)
	zeroLenSlice := make([]int, 0) // also empty, non-nil

	fmt.Printf("   nilSlice == nil: %v, len: %d\n", nilSlice == nil, len(nilSlice))
	fmt.Printf("   emptySlice == nil: %v, len: %d\n", emptySlice == nil, len(emptySlice))
	fmt.Printf("   zeroLenSlice == nil: %v, len: %d\n", zeroLenSlice == nil, len(zeroLenSlice))

	// All three have length 0, but nil slice is different!
	// nil slices are often used to represent "not allocated yet"

	// 4. FUNCTION RETURN VALUES
	fmt.Println("\n4. Function Return Values:")

	// Functions return zero values for uninitialized return variables
	result := exampleFunction()
	fmt.Printf("   Function returned: %v (zero value of int)\n", result)

	// 5. STRUCT INITIALIZATION
	fmt.Println("\n5. Struct Initialization Patterns:")

	// All these create a Person with zero values
	var p1 struct {
		Name string
		Age  int
	}
	p2 := struct {
		Name string
		Age  int
	}{}
	p3 := new(struct {
		Name string
		Age  int
	}) // Returns pointer to zero-valued struct

	fmt.Printf("   p1: %+v (zero values)\n", p1)
	fmt.Printf("   p2: %+v (zero values)\n", p2)
	fmt.Printf("   p3: %+v (pointer to zero-valued struct)\n", *p3)

	fmt.Println("\n=== KEY TAKEAWAYS ===")
	fmt.Println("1. Zero values make Go memory-safe: no undefined behavior")
	fmt.Println("2. Reference types are nil by default (must initialize before use)")
	fmt.Println("3. Aggregate types recursively initialize their elements")
	fmt.Println("4. Zero values enable simpler code with fewer explicit initializations")
	fmt.Println("5. The zero value is often a useful default (false, 0, \"\", nil)")
}

// Example function showing zero value return
func exampleFunction() int {
	var result int // Automatically initialized to 0
	// Some logic that might not assign to result
	return result // Returns 0 even if no assignment happened
}

/*
ADDITIONAL NOTES:

1. Why Zero Values?
   - Go aims for simplicity and safety
   - No "uninitialized variable" errors
   - Predictable program behavior
   - Less boilerplate initialization code

2. Common Pitfalls:
   - nil maps/slices: need to check before use
   - nil pointer dereference: always check pointers
   - Zero value might not be the desired initial state

3. Best Practices:
   - Leverage zero values when they make sense (counters start at 0)
   - Explicitly initialize when zero value isn't appropriate
   - Check for nil before using reference types
   - Use make() for maps, slices, channels when needed

4. Performance:
   - Zero initialization has minimal cost
   - Better than garbage values from uninitialized memory
   - Compiler optimizes zero-value initialization
*/
