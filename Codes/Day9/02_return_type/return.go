package main

import "fmt"

// ============================================================
// RETURN CONCEPTS IN GO
// ============================================================

// 1. BASIC RETURN
//    A function can return a value to the caller. The return type is declared after parameters.
//    Code after a return statement is unreachable and causes a compilation error.

func F4(x, y int) int {
	result := x + y
	fmt.Println("F4: result before return =", result)
	return result
	// fmt.Println("Hello World") // ❌ Unreachable code – compilation error
}

// 2. MULTIPLE RETURN STATEMENTS
//    A function can have multiple return statements, usually based on conditions.
//    Every code path must end with a return of the correct type.

func CheckNumber(a int) string {
	if a > 0 {
		return "Positive Number"
		// fmt.Println("This will never run") // ❌ Unreachable
	}
	if a < 0 {
		return "Negative Number"
	}
	// If neither condition is true, a must be zero
	return "Zero Number"
}

// 3. MULTIPLE RETURN VALUES
//    Go functions can return multiple values. This is often used to return a result and an error.
//    The return types are listed in parentheses: (type1, type2)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 4. NAMED RETURN VALUES (also called "naked return")
//    You can name the return values at the top of the function.
//    A bare return statement returns the current values of those named variables.
//    Named return values are initialized to their zero values.

func MyFunction(n string, a int) (Name string, Age int) {
	Age = a + 10        // Age is a named return variable
	Name = "Hello " + n // Name is a named return variable
	return              // naked return – returns Name and Age in order
	// return Name, Age // explicit return – also valid
}
func main() {
	// Calling F4
	fmt.Println("Calling F4(5,7):")
	resultF4 := F4(5, 7)
	fmt.Println("F4 returned:", resultF4)
	fmt.Println()

	// Calling CheckNumber
	fmt.Println("CheckNumber(10):", CheckNumber(10))
	fmt.Println("CheckNumber(-3):", CheckNumber(-3))
	fmt.Println("CheckNumber(0):", CheckNumber(0))
	fmt.Println()

	// Calling Divide (multiple return values)
	quotient, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Println("Divide(10,2) =", quotient)
	}

	quotient, err = Divide(5, 0)
	if err != nil {
		fmt.Println("Divide(5,0) error:", err)
	}
	fmt.Println()

	// Calling MyFunction (named returns)
	name, age := MyFunction("Arvinder", 34)
	fmt.Println("MyFunction returned:", name, age)
}

// ============================================================
// RETURN RULES – QUICK SUMMARY
// ============================================================
// • return immediately exits the function, returning given values.
// • Code after return is unreachable → compilation error.
// • Multiple return statements allowed; all paths must return correct type.
// • Multiple return values: (type1, type2, ...)
// • Named return values act as local variables; initialized to zero.
// • Naked return (just `return`) returns current values of named returns.
// • Named returns improve readability in short functions; use cautiously in long ones.
// • No function overloading: each function name must be unique in a package.
