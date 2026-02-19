package main

import "fmt"

// ============================================================
// DEMONSTRATING FUNCTION CALL ORDER AND CALL STACK
// ============================================================

// Example 1: Simple call order – functions execute in the order they are called.
func f1() {
	fmt.Println("  f1 called")
}

func f2() {
	fmt.Println("  f2 called")
}

func f3() {
	fmt.Println("  f3 called")
}

// Example 2:  calls – shows how the call stack builds and unwinds.
// Here, F3 calls F2, and F2 calls F1.
func F1() {
	fmt.Println("    F1 executes")
}

func F2() {
	fmt.Println("  F2 starts")
	F1() // calling F1 inside F2
	fmt.Println("  F2 ends")
}

func F3() {
	fmt.Println("F3 starts")
	F2() // calling F2 inside F3
	fmt.Println("F3 ends")
}

func main() {
	// ----------------------------------------
	// Example 1: Sequential calls
	// ----------------------------------------
	fmt.Println("=== Example 1: Simple call order ===")
	fmt.Println("main calls F1(), then F2(), then F3():")
	f1()
	f2()
	f3()
	// Output order: F1, F2, F3 (follows the call sequence)

	// ----------------------------------------
	// Example 2:  calls (call stack)
	// ----------------------------------------
	fmt.Println("\n=== Example 2:  calls (call stack) ===")
	fmt.Println("main calls F3():")
	F3()
	// Output:
	// F3 starts
	//   F2 starts
	//     F1 executes
	//   F2 ends
	// F3 ends

	// ----------------------------------------
	// Call Stack Visualization
	// ----------------------------------------
	/*
		STEP 1: main calls F3
		Stack (top is currently executing):
		+------------+
		| F3   |  <-- executing
		+------------+
		| main       |
		+------------+

		STEP 2: F3 calls F2
		Stack:
		+------------+
		| F2   |  <-- executing (F3 paused)
		+------------+
		| F3   |
		+------------+
		| main       |
		+------------+

		STEP 3: F2 calls F1
		Stack:
		+------------+
		| F1   |  <-- executing (F2 paused)
		+------------+
		| F2   |
		+------------+
		| F3   |
		+------------+
		| main       |
		+------------+

		STEP 4: F1 returns (pops from stack)
		Stack:
		+------------+
		| F2   |  <-- resumes execution
		+------------+
		| F3   |
		+------------+
		| main       |
		+------------+

		STEP 5: F2 returns (pops from stack)
		Stack:
		+------------+
		| F3   |  <-- resumes execution
		+------------+
		| main       |
		+------------+

		STEP 6: F3 returns (pops from stack)
		Stack:
		+------------+
		| main       |  <-- executing
		+------------+
	*/

	// Explanation:
	// When a function calls another, the current function is paused and pushed onto the call stack.
	// The new function is pushed on top and executes.
	// When it returns, it is popped, and the previous function resumes.
	// This LIFO (Last In, First Out) behavior is the call stack.
}

// ============================================================
// KEY TAKEAWAYS
// ============================================================
// • Functions execute only when called – order depends on call sequence, not definition order.
// •  calls create a stack: each called function is pushed, and returns are popped.
// • The call stack enables Go to remember where to return after a function completes.
