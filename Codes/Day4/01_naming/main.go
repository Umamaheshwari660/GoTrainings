//
// Go Naming Convention
//
// 1. Must begin with a letter (Unicode) or an underscore (_).
// 2. Can be followed by letters, digits, or underscores.
// 3. Identifiers are case-sensitive:
//    myVar and myvar are different identifiers.
// 4. CamelCase is the convention in Go:
//    - Exported identifiers (accessible outside the package) start with an uppercase letter.
//      Example: MyVar, MaxConnections
//    - Unexported (private) identifiers start with a lowercase letter.
//      Example: myVar, userName
// 5. Underscore (_) is commonly used as the blank identifier
//    to ignore values that are not needed.

package main

import "fmt"

func main() {
	// CamelCase for private (unexported) variables
	myName := "Arvinder"
	_ = myName
	coursesCompleted := 10
	_ = coursesCompleted

	// Uppercase for exported variables (typically at package scope)
	// These are exported because they start with a capital letter
	var MyPresence = 10
	_ = MyPresence

	// Using the blank identifier (_) to ignore an unused return value
	_, err := someFunction()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// someFunction demonstrates returning multiple values
func someFunction() (int, error) {
	return 42, nil
}
