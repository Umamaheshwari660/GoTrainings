// Switch Statement & Conditionals
package main

import "fmt"

func main() {
	fmt.Println("Q1. Vowel Checker")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	// - Handle both uppercase and lowercase letters
	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// - For non-alphabetic characters → "Not a letter"

	fmt.Println("Q2. BMI Categorization")
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	// Print the BMI value and category.

	fmt.Println("Q3. Ticket Price Calculator")
	// Calculate ticket price based on:
	// - Age (child: <12, adult: 12-64, senior: 65+)
	// - Day type (weekday/weekend)
	// - Student status (yes/no)

	// Rules:
	// - Base price: $10
	// - Children: 50% discount
	// - Seniors: 30% discount
	// - Students (adult age): 20% discount (only on weekdays)
	// - Weekend: +$2 extra

	fmt.Println("Q4. Roman Numeral to Integer")

	//  Convert Roman numerals to integers for values 1-10:
	// - I → 1
	// - II → 2
	// - III → 3
	// - IV → 4
	// - V → 5
	// - VI → 6
	// - VII → 7
	// - VIII → 8
	// - IX → 9
	// - X → 10
	// For invalid input → "Invalid Roman numeral"

	fmt.Println("Q5. Banking Transaction System")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.

	fmt.Println(" - - - - - - - - - - Theory Questions - - - - - - - - - - - ")

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present)
	//    d) Executes first case

	// 2. How does Go handle fallthrough in switch?
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough`
	//    d) No fallthrough allowed

	// 3. Which is NOT allowed in Go's if condition?
	//    a) if true { ... }
	//    b) if 1 { ... }
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }

	// 4. What is the scope of a variable declared in switch short statement?
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block
	//    d) Only the function

	// 5. Which is better for checking same variable against many values?
	//    a) Multiple if statements
	//    b) Switch statement
	//    c) Both are equal
	//    d) Ternary operator

	// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks
	// 2. T/F: Switch cases in Go are case-sensitive for strings
	// 3. T/F: You can use float values in switch cases
	// 4. T/F: Default case is mandatory in switch
	// 5. T/F: Tagless switch can have conditions in cases

}
