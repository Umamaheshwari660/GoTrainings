// Switch Statement & Conditionals
package main

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println("\n-----------Q1. Vowel Checker--------------")
	//   Write a program that takes a single character as input and determines if it's a vowel or consonant.
	// - Handle both uppercase and lowercase letters
	// - Use multiple values in cases (a, e, i, o, u, A, E, I, O, U)
	// - For non-alphabetic characters → "Not a letter"
     
	var ch string
	fmt.Print("Enter a character: ")
	fmt.Scan(&ch)

	switch ch {
	case "a", "e", "i", "o", "u",
		"A", "E", "I", "O", "U":
		fmt.Println("Vowel")

	case "b", "c", "d", "f", "g", "h", "j", "k",
		"l", "m", "n", "p", "q", "r", "s", "t",
		"v", "w", "x", "y", "z",
		"B", "C", "D", "F", "G", "H", "J", "K",
		"L", "M", "N", "P", "Q", "R", "S", "T",
		"V", "W", "X", "Y", "Z":
		fmt.Println("Consonant")

	default:
		fmt.Println("Not a letter")
	}



	fmt.Println("\n-----------Q2. BMI Categorization--------------")
	// Calculate BMI (Body Mass Index) using formula: BMI = weight(kg) / (height(m))²
	// Then categorize using tagless switch:
	// - BMI < 18.5 → "Underweight"
	// - 18.5 ≤ BMI < 25 → "Normal"
	// - 25 ≤ BMI < 30 → "Overweight"
	// - BMI ≥ 30 → "Obese"
	// Print the BMI value and category.

     
	var weight, height float64

	fmt.Print("Enter weight (kg): ")
	fmt.Scan(&weight)

	fmt.Print("Enter height (meters): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)

	fmt.Printf("BMI = %.2f\n", bmi)

	switch {
	case bmi < 18.5:
		fmt.Println("Underweight")
	case bmi < 25:
		
		fmt.Println("Normal")
	case bmi < 30:
		fmt.Println("Overweight")
	default:
		fmt.Println("Obese")
	}


	fmt.Println("\n-----------Q3. Ticket Price Calculator--------------")
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

    var age int
	var day, student string

	price := 10.0

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Print("Day (weekday/weekend): ")
	fmt.Scan(&day)

	fmt.Print("Student (yes/no): ")
	fmt.Scan(&student)

	// Age discount
	switch {
	case age < 12:
		price *= 0.5
	case age >= 65:
		price *= 0.7
	case age >= 12 && age <= 64 && student == "yes" && day == "weekday":
		price *= 0.8
	}

	// Weekend extra
	if day == "weekend" {
		price += 2
	}

	fmt.Println("Final Ticket Price: $", price)



	fmt.Println("\n-----------Q4. Roman Numeral to Integer")

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

    
	var roman string
	fmt.Print("Enter Roman numeral (I-X): ")
	fmt.Scan(&roman)

	roman = strings.ToUpper(roman)

	switch roman {
	case "I":
		fmt.Println(1)
	case "II":
		fmt.Println(2)
	case "III":
		fmt.Println(3)
	case "IV":
		fmt.Println(4)
	case "V":
		fmt.Println(5)
	case "VI":
		fmt.Println(6)
	case "VII":
		fmt.Println(7)
	case "VIII":
		fmt.Println(8)
	case "IX":
		fmt.Println(9)
	case "X":
		fmt.Println(10)
	default:
		fmt.Println("Invalid Roman numeral")
	}



	fmt.Println("\n-----------Q5. Banking Transaction System---------------")
	// Simulate banking operations:
	// - Input: transaction type ("deposit", "withdraw", "balance", "transfer")
	// - For deposit: add amount to balance
	// - For withdraw: check if sufficient balance
	// - For transfer: deduct from one account, add to another
	// - For invalid transaction → "Invalid operation"
	// Use switch with short statement to initialize variables.

    
	balance := 1000.0
	otherAccount := 500.0

	var operation string

	fmt.Print("Enter transaction (deposit/withdraw/balance/transfer): ")
	fmt.Scan(&operation)

	switch op := strings.ToLower(operation); op {

	case "deposit":
		var amt float64
		fmt.Print("Enter amount: ")
		fmt.Scan(&amt)
		balance += amt
		fmt.Println("New Balance:", balance)

	case "withdraw":
		var amt float64
		fmt.Print("Enter amount: ")
		fmt.Scan(&amt)

		if amt <= balance {
			balance -= amt
			fmt.Println("New Balance:", balance)
		} else {
			fmt.Println("Insufficient balance")
		}

	case "transfer":
		var amt float64
		fmt.Print("Enter amount: ")
		fmt.Scan(&amt)

		if amt <= balance {
			balance -= amt
			otherAccount += amt
			fmt.Println("Transfer successful")
			fmt.Println("Your Balance:", balance)
		} else {
			fmt.Println("Insufficient balance")
		}

	case "balance":
		fmt.Println("Current Balance:", balance)

	default:
		fmt.Println("Invalid operation")
	}


	// ====================================================
	// THEORY ANSWERS
	// ====================================================
	fmt.Println("\n--------- Theory Questions and Answers ---------")
     

	// Section 1: Multiple Choice Questions

	// 1. What happens when no case matches in a switch statement?
	//    a) Compilation error
	//    b) Runtime error
	//    c) Executes default case (if present)
	//    d) Executes first case

	fmt.Println("1 → c) Executes default case (if present)")

	// 2. How does Go handle fallthrough in switch?
	//    a) Automatic fallthrough by default
	//    b) Manual using `continue`
	//    c) Manual using `fallthrough`
	//    d) No fallthrough allowed

	fmt.Println("2 → c) Manual using `fallthrough`")

	// 3. Which is NOT allowed in Go's if condition?
	//    a) if true { ... }
	//    b) if 1 { ... }
	//    c) if x > 5 { ... }
	//    d) if condition() { ... }

	fmt.Println("3 → b) if 1 { ... } is not allowed in Go")

	// 4. What is the scope of a variable declared in switch short statement?
	//    a) Entire program
	//    b) Only the case where declared
	//    c) Entire switch block
	//    d) Only the function

	fmt.Println("4 → c) Entire switch block")

	// 5. Which is better for checking same variable against many values?
	//    a) Multiple if statements
	//    b) Switch statement
	//    c) Both are equal
	//    d) Ternary operator
	fmt.Println("5 → b) Switch statement")

    fmt.Println("\n--------- True/False Questions and Answers ---------")
// Section 2: True/False Questions

	// 1. T/F: Go requires braces {} for if blocks

	fmt.Println("1 → True")

	// 2. T/F: Switch cases in Go are case-sensitive for strings

	fmt.Println("2 → True")

	// 3. T/F: You can use float values in switch cases

    fmt.Println("3 → True")

	// 4. T/F: Default case is mandatory in switch
    fmt.Println("4 → False")


	// 5. T/F: Tagless switch can have conditions in cases
	fmt.Println("5 → True")

}
