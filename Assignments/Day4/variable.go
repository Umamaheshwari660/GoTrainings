package main

import (
	"fmt"
	"math"
	"unsafe"
	
)

//////////////////////////////////////////////////////
// 5. Functions (value return + pointer return)
//////////////////////////////////////////////////////

// returns value (stack)
func valueFunc() int {
	x := 10
	return x
}

// returns pointer (may escape to heap)
func pointerFunc() *int {
	y := 20
	return &y
}

//////////////////////////////////////////////////////
// 10. Grouped global variables
//////////////////////////////////////////////////////

var (
	g1 int    = 100
	g2 string // zero value ""
	g3 bool   // false
)

//////////////////////////////////////////////////////
// helper for Q9
//////////////////////////////////////////////////////

func multiReturn() (int, int) {
	return 5, 10
}

//////////////////////////////////////////////////////
// 12. Performance concept
//////////////////////////////////////////////////////

// dynamic append (slow)
func dynamicSlice() {
	var s []int
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
}

// pre-allocated (fast)
func preAllocatedSlice() {
	s := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
}

//////////////////////////////////////////////////////
// 13. Large array
//////////////////////////////////////////////////////

func largeArray() {
	arr := [1000000]int{}
	fmt.Println("Large array length:", len(arr))
}

//////////////////////////////////////////////////////

func main() {

	/*
		1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.

		2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.

		3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.

		4. Create two bool variables and print.

		5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.

		6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string.

		7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.

		8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.

		9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.

		10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.

		11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.

		12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).

		13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".

		14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.

		15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.

		16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.

		17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.

		18. Redeclare with one new variable: x, z := 50, 60. Print both.

		19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.

		20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types.


	*/
     
	//////////////////////////////////////////////////
	// 1. Integer max values
	// 1. Declare variables of types int8, int16, int32, int64, uint8, uint16, uint32, and uint64. Assign them their maximum values and print.
	//////////////////////////////////////////////////
	var i8 int8 = math.MaxInt8
	var i16 int16 = math.MaxInt16
	var i32 int32 = math.MaxInt32
	var i64 int64 = math.MaxInt64

	var u8 uint8 = math.MaxUint8
	var u16 uint16 = math.MaxUint16
	var u32 uint32 = math.MaxUint32
	var u64 uint64 = math.MaxUint64

	fmt.Println("1.", i8, i16, i32, i64, u8, u16, u32, u64)

	//////////////////////////////////////////////////
	// 2. Overflow
	// 2. Demonstrate integer overflow by assigning uint8 variable with 255 then adding 1. Print before and after values.
	//////////////////////////////////////////////////
	var x uint8 = 255
	fmt.Println("2. Before:", x)
	x++
	fmt.Println("   After :", x) // becomes 0

	//////////////////////////////////////////////////
	// 3. float precision
	// 3. Show precision differences between float32 and float64 by computing 1.0/3.0 with both types and printing with 10 decimal places.
	//////////////////////////////////////////////////
	f32 := float32(1.0 / 3.0)
	f64 := float64(1.0 / 3.0)

	fmt.Printf("3. float32: %.10f\n", f32)
	fmt.Printf("   float64: %.10f\n", f64)

	//////////////////////////////////////////////////
	// 4. bool
	// 4. Create two bool variables and print.
	//////////////////////////////////////////////////
	b1 := true
	b2 := false
	fmt.Println("4.", b1, b2)

	//////////////////////////////////////////////////
	// 5. functions	
    // 5. Write two functions - one returning a value, one returning a pointer to local variable. Build with -gcflags="-m" to analyze escape.
	//////////////////////////////////////////////////
	fmt.Println("5.", valueFunc(), *pointerFunc())

	//////////////////////////////////////////////////
	// 6. unsafe.Sizeof
	// 6. Use unsafe.Sizeof() to print sizes of all integer types, float types, bool, and string
	//////////////////////////////////////////////////
	fmt.Println("6. Sizes:")
	fmt.Println("int8:", unsafe.Sizeof(i8))
	fmt.Println("int16:", unsafe.Sizeof(i16))
	fmt.Println("int32:", unsafe.Sizeof(i32))
	fmt.Println("int64:", unsafe.Sizeof(i64))
	fmt.Println("float32:", unsafe.Sizeof(f32))
	fmt.Println("float64:", unsafe.Sizeof(f64))
	fmt.Println("bool:", unsafe.Sizeof(b1))
	fmt.Println("string:", unsafe.Sizeof("hello"))

	//////////////////////////////////////////////////
	// 7. Exported/unexported (concept)
	// 7. Create exported and unexported variables in a separate package. Try accessing both from main package and handle the error.
	//////////////////////////////////////////////////
	fmt.Println("7. Use CapitalName to export from other package")

	//////////////////////////////////////////////////
	// 8. Shadowing
	// 8. Declare x := 10 in main. Inside a block, redeclare x := 20. Print both values to demonstrate shadowing.
	//////////////////////////////////////////////////
	x2 := 10
	{
		x2 := 20
		fmt.Println("8. inner:", x2)
	}
	fmt.Println("   outer:", x2)

	//////////////////////////////////////////////////
	// 9. blank identifier
	// 9. Use _ to ignore multiple return values from a function. Try to declare _ as a variable and assign to it.
	//////////////////////////////////////////////////
	a, _ := multiReturn()
	fmt.Println("9.", a)
	_ = 100 // allowed

	//////////////////////////////////////////////////
	// 10. global vars
	// 10. Create package-level variables using grouped declaration at global scope. Initialize some, leave others with zero values.
	//////////////////////////////////////////////////
	fmt.Println("10.", g1, g2, g3)
	fmt.Println("10.", g1, g2, g3)

	//////////////////////////////////////////////////
	// 11. for loop variable
	// 11. Declare variable inside for loop: for i := 0; i < 3; i++ { x := i*2 }. Try to access x outside loop and handle error.
	//////////////////////////////////////////////////
	for i := 0; i < 3; i++ {
		x := i * 2
		fmt.Println("11 inside:", x)
	}
	// cannot use x outside -> compile error

	//////////////////////////////////////////////////
	// 12. performance
	// 12. Compare performance concept: Create slice with pre-allocated capacity vs dynamic appends (just write both versions).
	//////////////////////////////////////////////////
	dynamicSlice()
	preAllocatedSlice()
	fmt.Println("12. slices created")

	//////////////////////////////////////////////////
	// 13. large array
	// 13. Create large array [1000000]int locally in function. Check if it escapes with -gcflags="-m".
	//////////////////////////////////////////////////
	largeArray()

	//////////////////////////////////////////////////
	// 14. zero values
	// 14. Declare variables without initialization: var a int; var b float64; var c string; var d bool; var e *int; var f []int; var g map[string]int. Print them to show zero values.
	//////////////////////////////////////////////////
	var a1 int
	var b float64
	var c string
	var d bool
	var e *int
	var f []int
	var g map[string]int

	fmt.Println("14.", a1, b, c, d, e, f, g)

	//////////////////////////////////////////////////
	// 15. multiple declaration
	// 15. Declare multiple variables in single line: var a, b, c int = 1, 2, 3. Then update them in single assignment: a, b, c = 4, 5, 6.
	//////////////////////////////////////////////////
	var p, q, r int = 1, 2, 3
	fmt.Println("15 before:", p, q, r)
	p, q, r = 4, 5, 6
	fmt.Println("   after :", p, q, r)

	//////////////////////////////////////////////////
	// 16. swap
	// 16. Swap two variables without temporary: x, y := 10, 20; x, y = y, x. Print before and after.
	//////////////////////////////////////////////////
	m, n := 10, 20
	fmt.Println("16 before:", m, n)
	m, n = n, m
	fmt.Println("   after :", m, n)

	//////////////////////////////////////////////////
	// 17. redeclare error (commented)
	// 17. Try redeclaring variables with := when no new variables: x, y := 10, 20; x, y := 30, 40. Handle error.
	//////////////////////////////////////////////////
	// x, y := 10, 20
	// x, y := 30, 40 // ERROR

	//////////////////////////////////////////////////
	// 18. one new variable
	// 18. Redeclare with one new variable: x, z := 50, 60. Print both.
	//////////////////////////////////////////////////
	x3 := 50
	x3, z := 60, 70
	fmt.Println("18.", x3, z)

	//////////////////////////////////////////////////
	// 19. math constants
	// 19. Use math.MaxInt32, math.MinInt32, math.MaxUint32 constants. Assign them to variables and print.
	//////////////////////////////////////////////////
	fmt.Println("19.",
		math.MaxInt32,
		math.MinInt32,
		math.MaxUint32,
	)

	//////////////////////////////////////////////////
	// 20. nil check
	// 20. Declare variables with zero values and check if they equal nil for pointer, slice, and map types. these can make the simple mothod you can write 
	//////////////////////////////////////////////////
	var p1 *int
	var s1 []int
	var m1 map[string]int

	fmt.Println("20.", p1 == nil, s1 == nil, m1 == nil)


}
