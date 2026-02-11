package main
import "fmt"



//14 Grouped declaration
var (
	g1 = "Global grouped 1"
	g2 = "Global grouped 2"
)

// Q12 Grouped declarations
var (
	salary = 10000
	age    = 30
	id     = 12
)
// Q11 Individual declarations
var salary1 int = 10000
var age1 int = 30
var id1 int = 12

// Q10 Function (local scope )
func varUse() {
	local := 500
	fmt.Println("Q10 Local variable:", local)
}


/*
=====================================================
Day 3 – Golang Training - Coding Questions
=====================================================
*/

/*
-----------------------------------------------------
1. BASIC VARIABLE DECLARATION
-----------------------------------------------------
*/

// Q1. Declare an integer variable x with value 200 and print it.

// Q2. Declare a variable y using type inference with value 4000 and print it.

// Q3. Declare a variable a of type int, initialize it later with value 300, and print it.

// Q4. Use short declaration to create a string variable slogan with value
//     "Variables in Go - Day 3" and print it.

/*
-----------------------------------------------------
2. STATIC TYPING
-----------------------------------------------------
*/

// Q5. Declare a variable x as int and assign it a number.
//     Then try assigning a string to x and observe the compiler error.

// Q6. Declare a variable using := with a string value.
//     Then try assigning an integer to the same variable.

/*
-----------------------------------------------------
3. VARIABLE SCOPE
-----------------------------------------------------
*/

// Q7. Declare a variable x inside main() and print it.

// Q8. Create a block {} inside main().
//     Declare a variable y inside the block and print it inside the block.
//     Try printing y outside the block.

// Q9. Declare a global variable salary = 10000.
//     Access and print it inside main().

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().

/*
-----------------------------------------------------
4. GLOBAL AND GROUPED DECLARATIONS
-----------------------------------------------------
*/

// Q11. Declare global variables age, salary, and id using individual declarations.

// Q12. Rewrite Q11 using grouped variable declaration.

// Q13. Declare grouped variables inside main() and print them.

// Q14. Write code to demonstrate that grouped declarations work
//      both at global scope and function scope.

/*
-----------------------------------------------------
5. MULTIPLE DECLARATION AND ASSIGNMENT
-----------------------------------------------------
*/

// Q15. Declare two variables b and c in a single line with values 10 and 20.
//      Print both variables.

// Q16. Update variables b and c using a single assignment statement.

// Q17. Swap two variables using multiple assignment
//      without using a temporary variable.

// Q18. Declare variables x and y using :=
//      Then try redeclaring x and y again using := and observe the error.

// Q19. Redeclare x using := by introducing one new variable.
//      Print both variables.

/*
-----------------------------------------------------
6. ZERO VALUES
-----------------------------------------------------
*/

// Q20. Declare variables of type int, float64, string, bool,
//      slice, and pointer without initialization.
//      Print their values and check which ones are nil.

func main() {



	// Q1
	var x int = 200
	fmt.Println("Q1:", x)

//  Q2
	var y = 4000
	fmt.Println("Q2:", y)

// Q3
	var a int
	a = 300
	fmt.Println("Q3:", a)

// Q4
	slogan := "Variables in Go "
	fmt.Println("Q4:", slogan)



/*
-----------------------------------------------------
2. STATIC TYPING
-----------------------------------------------------
*/

// Q5. Declare a variable x as int and assign it a number.
//     Then try assigning a string to x and observe the compiler error.

// Q6. Declare a variable using := with a string value.
//     Then try assigning an integer to the same variable.




// Q5
	var num int = 10
	fmt.Println("Q5:", num)

	// num = "hello" ❌ ERROR (cannot assign string to int)

	
// Q6
	text := "Go Language"
	fmt.Println("Q6:", text)

	// text = 100 ❌ ERROR (cannot assign int to string)




/*
-----------------------------------------------------
3. VARIABLE SCOPE
-----------------------------------------------------
*/

// Q7. Declare a variable x inside main() and print it.

// Q8. Create a block {} inside main().
//     Declare a variable y inside the block and print it inside the block.
//     Try printing y outside the block.

// Q9. Declare a global variable salary = 10000.
//     Access and print it inside main().

// Q10. Create a function varUse() that declares and prints a local variable.
//      Try accessing that variable from main().

// Q7
	x2 := 50
	fmt.Println("Q7:", x2)

// Q8
	{
		y2 := 99
		fmt.Println("Q8:","Inside block", y2)
	}
// fmt.Println(y2) ❌ ERROR (block scope)

// Q9
	fmt.Println("Q9 Salary:", salary)

// Q10
	varUse()
	// fmt.Println(local) ❌ ERROR (not accessible)



/*
-----------------------------------------------------
4. GLOBAL AND GROUPED DECLARATIONS
-----------------------------------------------------
*/

// Q11. Declare global variables age, salary, and id using individual declarations.

// Q12. Rewrite Q11 using grouped variable declaration.

// Q13. Declare grouped variables inside main() and print them.

// Q14. Write code to demonstrate that grouped declarations work
//      both at global scope and function scope.

// Q11
fmt.Println("Q11 Individual:", salary1, age1, id1)
// Q12
fmt.Println("Q12 Grouped:", salary, age, id)


// Q13 - grouped inside function
	var (
		localAge    = 25
		localSalary = 50000
	)

	fmt.Println("Q13:", localAge, localSalary)

	// Q14
	fmt.Println("Q14:", g1, g2)


/*
-----------------------------------------------------
5. MULTIPLE DECLARATION AND ASSIGNMENT
-----------------------------------------------------
*/

// Q15. Declare two variables b and c in a single line with values 10 and 20.
//      Print both variables.

// Q16. Update variables b and c using a single assignment statement.

// Q17. Swap two variables using multiple assignment
//      without using a temporary variable.

// Q18. Declare variables x and y using :=
//      Then try redeclaring x and y again using := and observe the error.

// Q19. Redeclare x using := by introducing one new variable.
//      Print both variables.



	// Q15
	var b, c = 10, 20
	fmt.Println("Q15:", b, c)

	// Q16
	b, c = 30, 40
	fmt.Println("Q16 Updated:", b, c)

	// Q17 (swap)
	// b, c = c, b
	// fmt.Println("Q17 Swapped:", b, c)

	b, c = 10, 20
	fmt.Println("Q17 Before Swap:", b, c)
	// Step 2: swap using multiple assignment
	b, c = c, b
	fmt.Println("Q17 After Swap:", b, c)

	// Q18
	p, q := 5, 6
	fmt.Println("Q18:", p, q)
	// p, q := 7, 8 ❌ ERROR (no new variable)

	// Q19
	p, newVar := 100, 200
	fmt.Println("Q19:", p, newVar)


/*
-----------------------------------------------------
6. ZERO VALUES
-----------------------------------------------------
*/

// Q20. Declare variables of type int, float64, string, bool,
//      slice, and pointer without initialization.
//      Print their values and check which ones are nil.

    fmt.Println("Q20:")
// int
	var zi int
	fmt.Println("int:", zi)

	// float64
	var zf float64
	fmt.Println("float64:", zf)

	// string
	var zs string
	fmt.Println("string:", zs)

	// bool
	var zb bool
	fmt.Println("bool:", zb)

	// slice
	var zsli []int
	fmt.Println("slice:", zsli)
	fmt.Println("slice is nil:", zsli == nil)

	// pointer
	var zp *int
	fmt.Println("pointer:", zp)
	fmt.Println("pointer is nil:", zp == nil)
}

