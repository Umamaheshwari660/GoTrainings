package main

import "fmt"

func main() {
	fmt.Println("Comparison Operators in Go")
	fmt.Println("Comparison Operators -- Returns True or False")

	a := 100
	b := 200
	c := 100
	fmt.Println("a =", a, "b =", b, "c =", c)
	fmt.Println()

	fmt.Println("a == b:", a == b)
	fmt.Println("a == c:", a == c)
	fmt.Println("a != b:", a != b)
	fmt.Println("a != c:", a != c)
	fmt.Println("a > b:", a > b)
	fmt.Println("a > c:", a > c)
	fmt.Println("a < b:", a < b)
	fmt.Println("a < c:", a < c)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a >= c:", a >= c)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a <= c:", a <= c)

	fmt.Println("\nComparing strings:")
	s1 := "hello"
	s2 := "World"
	s3 := "Hello"
	fmt.Println("s1 =", s1, "s2 =", s2, "s3 =", s3)
	fmt.Println("s1 == s2:", s1 == s2)
	fmt.Println("s2 == s3:", s2 == s3)
	fmt.Println("s3 == s1:", s3 == s1)
	fmt.Println("s1 == s3:", s1 == s3)

	// Note: String comparison in Go is case-sensitive
	// and compares Unicode code points
}
