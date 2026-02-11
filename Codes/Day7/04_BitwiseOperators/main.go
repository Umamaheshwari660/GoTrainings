package main

import "fmt"

func main() {
	fmt.Println("BITWISE OPERATORS IN GO")
	fmt.Println("=======================")

	// 1. BITWISE AND (&)
	fmt.Println("\n1. BITWISE AND (&)")
	fmt.Println("   Truth Table:")
	fmt.Println("   0 & 0 = 0")
	fmt.Println("   0 & 1 = 0")
	fmt.Println("   1 & 0 = 0")
	fmt.Println("   1 & 1 = 1")

	fmt.Println("\n   Examples:")
	fmt.Println("   5 & 3 =", 5&3)   // 0101 & 0011 = 0001 = 1
	fmt.Println("   10 & 6 =", 10&6) // 1010 & 0110 = 0010 = 2
	fmt.Println("   15 & 7 =", 15&7) // 1111 & 0111 = 0111 = 7

	// 2. BITWISE OR (|)
	fmt.Println("\n\n2. BITWISE OR (|)")
	fmt.Println("   Truth Table:")
	fmt.Println("   0 | 0 = 0")
	fmt.Println("   0 | 1 = 1")
	fmt.Println("   1 | 0 = 1")
	fmt.Println("   1 | 1 = 1")

	fmt.Println("\n   Examples:")
	fmt.Println("   5 | 3 =", 5|3)   // 0101 | 0011 = 0111 = 7
	fmt.Println("   10 | 6 =", 10|6) // 1010 | 0110 = 1110 = 14
	fmt.Println("   12 | 3 =", 12|3) // 1100 | 0011 = 1111 = 15

	// 3. BITWISE XOR (^)
	fmt.Println("\n\n3. BITWISE XOR (^)")
	fmt.Println("   Truth Table:")
	fmt.Println("   0 ^ 0 = 0")
	fmt.Println("   0 ^ 1 = 1")
	fmt.Println("   1 ^ 0 = 1")
	fmt.Println("   1 ^ 1 = 0")

	fmt.Println("\n   Examples:")
	fmt.Println("   5 ^ 3 =", 5^3)   // 0101 ^ 0011 = 0110 = 6
	fmt.Println("   10 ^ 6 =", 10^6) // 1010 ^ 0110 = 1100 = 12
	fmt.Println("   9 ^ 5 =", 9^5)   // 1001 ^ 0101 = 1100 = 12

	// 4. BITWISE NOT (^)
	fmt.Println("\n\n4. BITWISE NOT (^)")
	fmt.Println("   Truth Table:")
	fmt.Println("   ^0 = 1")
	fmt.Println("   ^1 = 0")

	fmt.Println("\n   Examples:")
	fmt.Println("   ^5 =", ^5) // ^0101 = 1010 = -6 (in 4-bit signed)
	fmt.Println("   ^3 =", ^3) // ^0011 = 1100 = -4 (in 4-bit signed)
	fmt.Println("   ^0 =", ^0) // ^0000 = 1111 = -1 (in 4-bit signed)

	// Note: ^ works differently in Go for signed integers
	// It does two's complement

	// 5. BITWISE AND NOT (&^)
	fmt.Println("\n\n5. BITWISE AND NOT (&^)")
	fmt.Println("   Truth Table:")
	fmt.Println("   0 &^ 0 = 0")
	fmt.Println("   0 &^ 1 = 0")
	fmt.Println("   1 &^ 0 = 1")
	fmt.Println("   1 &^ 1 = 0")

	fmt.Println("\n   Examples:")
	fmt.Println("   5 &^ 3 =", 5&^3)   // 0101 &^ 0011 = 0100 = 4
	fmt.Println("   10 &^ 6 =", 10&^6) // 1010 &^ 0110 = 1000 = 8
	fmt.Println("   15 &^ 5 =", 15&^5) // 1111 &^ 0101 = 1010 = 10

	// 6. LEFT SHIFT (<<)
	fmt.Println("\n\n6. LEFT SHIFT (<<)")
	fmt.Println("   Moves bits to the left")
	fmt.Println("   Fills right with zeros")

	fmt.Println("\n   Examples:")
	fmt.Println("   1 << 0 =", 1<<0) // 0001 = 1
	fmt.Println("   1 << 1 =", 1<<1) // 0010 = 2
	fmt.Println("   1 << 2 =", 1<<2) // 0100 = 4
	fmt.Println("   1 << 3 =", 1<<3) // 1000 = 8

	fmt.Println("\n   Multiply by 2 examples:")
	fmt.Println("   5 << 1 =", 5<<1) // 5 * 2 = 10
	fmt.Println("   5 << 2 =", 5<<2) // 5 * 4 = 20
	fmt.Println("   5 << 3 =", 5<<3) // 5 * 8 = 40

	// 7. RIGHT SHIFT (>>)
	fmt.Println("\n\n7. RIGHT SHIFT (>>)")
	fmt.Println("   Moves bits to the right")
	fmt.Println("   For positive numbers, fills left with zeros")

	fmt.Println("\n   Examples:")
	fmt.Println("   8 >> 0 =", 8>>0) // 1000 = 8
	fmt.Println("   8 >> 1 =", 8>>1) // 0100 = 4
	fmt.Println("   8 >> 2 =", 8>>2) // 0010 = 2
	fmt.Println("   8 >> 3 =", 8>>3) // 0001 = 1

	fmt.Println("\n   Divide by 2 examples:")
	fmt.Println("   20 >> 1 =", 20>>1) // 20 / 2 = 10
	fmt.Println("   20 >> 2 =", 20>>2) // 20 / 4 = 5
	fmt.Println("   20 >> 3 =", 20>>3) // 20 / 8 = 2

	// PRACTICAL EXAMPLES
	fmt.Println("\n\nPRACTICAL EXAMPLES")
	fmt.Println("=================")

	// Example 1: Check if number is even or odd
	fmt.Println("\n1. Check if number is even or odd:")
	num1 := 7
	num2 := 10
	fmt.Println("   Number", num1, "is odd:", (num1&1) == 1)
	fmt.Println("   Number", num2, "is even:", (num2&1) == 0)

	// Example 2: Multiply and divide by powers of 2
	fmt.Println("\n2. Multiply and divide by powers of 2:")
	x := 25
	fmt.Println("   Number:", x)
	fmt.Println("   Multiply by 2:", x<<1) // 50
	fmt.Println("   Multiply by 4:", x<<2) // 100
	fmt.Println("   Divide by 2:", x>>1)   // 12
	fmt.Println("   Divide by 4:", x>>2)   // 6

	// Example 3: Swapping two numbers
	fmt.Println("\n3. Swap two numbers using XOR:")
	a := 5
	b := 9
	fmt.Println("   Before swap: a =", a, "b =", b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("   After swap: a =", a, "b =", b)

	// Example 4: Create permissions
	fmt.Println("\n4. Create permission flags:")
	read := 1    // 001
	write := 2   // 010
	execute := 4 // 100

	user1 := read | execute         // 101 = can read and execute
	user2 := read | write           // 011 = can read and write
	user3 := read | write | execute // 111 = can do everything

	fmt.Println("   User1 permission value:", user1)
	fmt.Println("   User2 permission value:", user2)
	fmt.Println("   User3 permission value:", user3)

	// Check if user1 can read
	fmt.Println("   User1 can read:", (user1&read) != 0)
	fmt.Println("   User1 can write:", (user1&write) != 0)

	fmt.Println("\n=== END ===")
}
