// FUNCTIONS IN GO
// ============================================================================
//  WHAT IS A FUNCTION?
// ============================================================================
//
/*
Definition:
-----------
A function is a reusable block of code that performs a specific task.
It is defined once and can be executed (called) many times from different
parts of the program.

Why use functions?
------------------
1. Reusability – Write once, use many times.
2. Modularity – Break complex problems into smaller, manageable pieces.
3. Readability – Named operations make code self-documenting.
4. Maintainability – Fix or improve a function in one place.
5. Testing – Each function can be tested independently.

Important Execution Rules:
--------------------------
1. Functions will execute only if they are called.
   Defining a function does not run it; you must invoke it explicitly.

2. The order of execution of functions depends upon the order in which they are called.
   Unlike the sequential order of declaration, the actual runtime order is determined
   by the call sequence in the calling function (e.g., main).

3. Code written after 'return' will not execute, and Go will throw a compilation error.
   This means 'return' also marks the end of the function block – any statement
   placed after it is unreachable.

Characteristics:
----------------
- A function executes only when it is called/invoked.
- It may receive input values (parameters) and may return one or more results.
- The function body is enclosed in curly braces { }.
- In Go, functions are first-class citizens – they can be assigned to variables,
  passed as arguments, and returned from other functions.

============================================================================
TOPIC 2: FUNCTION SYNTAX
============================================================================

Basic Syntax:
-------------
  func functionName(parameterList) returnType {
      // body
  }

Components:
-----------
1. func        : Keyword that declares a function.
2. functionName : Identifier; must follow Go naming rules.
3. parameterList: Comma-separated list of input parameters (name type).
                  Can be empty: ().
4. returnType   : Data type of the value returned. If nothing is returned,
                  this part is omitted entirely.
5. Body         : Block of statements inside { }.

Parameter Shortcut:
-------------------
If multiple consecutive parameters share the same type, you can write:
  func add(x, y int) int       // instead of (x int, y int)

Return Statement:
-----------------
- The 'return' keyword sends a value back to the caller.
- A function can have multiple return statements (e.g., early returns).
- If the function declares a return type, every code path must end with return.
- IMPORTANT: Any code written after a return statement is unreachable and
  will cause a compilation error: "unreachable code".

Function Without Parameters, Without Return:
--------------------------------------------
  func sayHello() {
      fmt.Println("Hello!")
  }

Function With Parameters, Without Return:
-----------------------------------------
  func greet(name string) {
      fmt.Println("Hello", name)
  }

Function Without Parameters, With Return:
-----------------------------------------
  func getAnswer() int {
      return 42
  }

Function With Parameters, With Return:
--------------------------------------
  func multiply(a, b int) int {
      return a * b
  }
*/
// ============================================================================
// CODE EXAMPLES – FOUR TYPES OF USER-DEFINED FUNCTIONS
// ============================================================================

package main

import "fmt"

// ------------------------------------------------------------
// Type 1: WITHOUT PARAMETERS, WITHOUT RETURN
// ------------------------------------------------------------
func F1() {
	fmt.Println("F1: No parameters, no return – just prints a message.")
	// Code after return would cause error:
	// return
	// fmt.Println("This will never execute") // ❌ compilation error
}

// ------------------------------------------------------------
// Type 2: WITH PARAMETERS, WITHOUT RETURN
// ------------------------------------------------------------
func F2(x, y int) {
	fmt.Println("F2: Parameters only – sum =", x+y)
	// This function has no 'return' because no return type is declared.
}

// ------------------------------------------------------------
// Type 3: WITHOUT PARAMETERS, WITH RETURN
// ------------------------------------------------------------
func F3() int {
	x := 10
	y := 20
	return x + y
	// fmt.Println("Done") // ❌ unreachable code – compilation error
}

// ------------------------------------------------------------
// Type 4: WITH PARAMETERS, WITH RETURN
// ------------------------------------------------------------
func F4(x, y int) int {
	return x + y
	// fmt.Println("This is after return") // ❌ unreachable code
}

// ------------------------------------------------------------
// SYSTEM DEFINED FUNCTION (ENTRY POINT)
// ------------------------------------------------------------
func main() {
	// Execution order is determined by the sequence of calls, not by declaration order.

	// 1st call
	F1()

	// 2nd call
	a, b := 10, 20
	F2(a, b)

	// 3rd call
	fmt.Println("F3 result:", F3())
	x := F3()
	fmt.Println("F3 result + 10:", x+10)

	// 4th call
	fmt.Println("F4(20,50) =", F4(20, 50))
	t := F4(40, 50)
	fmt.Println("F4 result stored:", t)

	// Even though F3 is defined above F4 in the file, it executes only when called here.
	// The order of execution follows the order of calls inside main.
}
