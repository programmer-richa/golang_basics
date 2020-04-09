/*
Data types specify the type of data that a valid Go variable can hold.
In Go language, the type is divided into four categories which are as follows:

Basic type: Numbers, strings, and booleans come under this category.
Aggregate type: Array and structs come under this category.
Reference type: Pointers, slices, maps, functions, and channels come under this category.
Interface type

A Variable is a placeholder of the information which can be changed at runtime.
And variables allow to Retrieve and Manipulate the stored information.

Using var keyword: In Go language, variables are created using var keyword of a particular type,
connected with name and provide its initial value.

Syntax:

var variable_name type = expression

Note: If the expression is removed, then the variable holds zero-value for the type like zero for the number, false for Booleans,
“” for strings, and nil for interface and reference type. So, there is no such concept of an uninitialized variable in Go language.


Using short variable declaration: The local variables which are declared and initialize in the functions are declared by using short variable declaration.
Syntax:

variable_name:= expression
*/
package variables

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

// Variable Modules Listing Constants
const (
	_ = iota
	declarationConst
	typeConst
	zeroConst
	sameTypeConst
	differentTypeConst
	functionConst
	blankConst
	exitConst
	declarationConstLbl   = "Variable declaration and initialization "
	typeConstLbl          = "Variable Type Tracking"
	zeroConstLbl          = "Zero Values Of Basic Data Types"
	sameTypeConstLbl      = "Declaring Multiple Variables With Same Data Type In a Single Line"
	differentTypeConstLbl = "Declaring Multiple Variables With Different Data Type In a Single Line"
	functionConstLbl      = "Initializing Variables From a Function Returning Multiple Values"
	blankConstLbl         = "Using Blank Identifier to Overcome The Compultion to Use Variable In Code"
	exitConstLbl          = "Exit"
)

//Block Works according to choice of user
func Block() {

	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, declarationConst, declarationConstLbl)
		utility.Println(level, typeConst, typeConstLbl)
		utility.Println(level, zeroConst, zeroConstLbl)
		utility.Println(level, sameTypeConst, sameTypeConstLbl)
		utility.Println(level, differentTypeConst, differentTypeConstLbl)
		utility.Println(level, functionConst, functionConstLbl)
		utility.Println(level, blankConst, blankConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case declarationConst:
			Declaration()
		case typeConst:
			Type()
		case zeroConst:
			Zero()
		case sameTypeConst:
			SameType()
		case differentTypeConst:
			DifferentType()
		case functionConst:
			Function()
		case blankConst:
			Blank()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Declaration It demonstrates alternatives to declare and initialize variables
func Declaration() {
	utility.Printh(sublevel, declarationConstLbl)
	// declared variables must be used at least once in the program, otherwise compiler generated CTE
	// Way 1: Declare and initialize variables explicitly
	var i int // declaration of int type variable i
	i = 10    // assigning value to int
	utility.Println(sublevel, "Integer value assigned to variable i = ", i)

	// Way 2: Declare and initialize variables explicitly on the same line
	var f float32 = 25.46 // declaration of float type variable f for 32 bit memory
	utility.Println(sublevel, "Float (32- bit )value assigned to variable f = ", f)

	var f1 float64 // declaration of float type variable f for 64 bit memory
	f1 = 25.46
	utility.Println(sublevel, "Float (64- bit )value assigned to variable f = ", f1)

	//Way 3: Implicit variable declaration, here compiler will interpret the data type based on the value assigned to the variable
	// This is most frequent way of declaring variables
	firstName := "Richa"
	utility.Println(sublevel, "First Name : "+firstName)

	var b bool = true // Boolean Data Type
	utility.Println(sublevel, "Boolean value : ", b)

	var realNumber float32 = 5
	var imaginaryNumber float32 = 4
	// complex function expects arguments to be float
	complexNumber := complex(realNumber, imaginaryNumber)
	utility.Println(level, "Complex Number : ", complexNumber)

	// retrieving real and imaginary part of a complex number
	utility.Println(level, "Real value  : ", real(complexNumber))
	utility.Println(level, "Imaginary value  : ", imag(complexNumber))
}

//Type It demonstrates value and type tracking of a variable
func Type() {
	// Variable declared and
	// initialized without the
	// explicit type
	utility.Printh(sublevel, typeConstLbl)
	var myvariable1 = 20
	var myvariable2 = "GeeksforGeeks"
	var myvariable3 = 34.80
	tabs := utility.Tabs(sublevel)
	// Display the value and the
	// type of the variables
	fmt.Printf(tabs+"The value of myvariable1 is : %d\n", myvariable1)
	fmt.Printf(tabs+"The type of myvariable1 is : %T\n", myvariable1)
	fmt.Printf(tabs+"The value of myvariable2 is : %s\n", myvariable2)
	fmt.Printf(tabs+"The type of myvariable2 is : %T\n", myvariable2)
	fmt.Printf(tabs+"The value of myvariable3 is : %f\n", myvariable3)
	fmt.Printf(tabs+"The type of myvariable3 is : %T\n", myvariable3)
}

//Zero  Displays default values of basic data types
func Zero() {
	// Variable declared and
	// initialized without expression
	tabs := utility.Tabs(sublevel)
	utility.Printh(sublevel, zeroConstLbl)
	var int1 int
	var int2 int8
	var int3 int16
	var int4 int32
	var int5 int64

	var uint1 uint
	var uint2 uint8
	var uint3 uint16
	var uint4 uint32
	var uint5 uint64

	var float1 float32
	var float2 float64

	var complex1 complex64
	var complex2 complex128

	var string1 string

	var bool1 bool

	// Display the zero-value of the variables
	// utility.PrintLevelString(level,The value of %T is : %d\n", int1, int1)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, int1)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int2, int2)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int3, int3)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int4, int4)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int5, int5)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, int1)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, uint1)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, uint2)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, uint3)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, uint4)
	fmt.Printf(tabs+"The zero value of %T is : %d\n", int1, uint5)
	fmt.Printf(tabs+"The zero value of %T is : %f\n", int1, float1)
	fmt.Printf(tabs+"The zero value of %T is : %f\n", int1, float2)
	fmt.Printf(tabs+"The zero value of %T is : %s\n", string1, string1)
	fmt.Printf(tabs+"The zero value of %T is : %t\n", bool1, bool1)
	fmt.Printf(tabs+"The zero value of %T is : %g\n", complex1, complex1)
	fmt.Printf(tabs+"The zero value of %T is : %g\n", complex2, complex2)

}

//SameType  Demonstrates how multiple variables having same data type can be declared on a single line
func SameType() {
	// Multiple variables of the same type
	// are declared and initialized
	// in the single line
	tabs := utility.Tabs(sublevel)
	utility.Printh(sublevel, sameTypeConstLbl)
	var myvariable1, myvariable2, myvariable3 int = 2, 454, 67

	// Display the values of the variables
	fmt.Printf(tabs+"myvariable1 %d\n", myvariable1)
	fmt.Printf(tabs+"myvariable2 %d\n", myvariable2)
	fmt.Printf(tabs+"myvariable3 %d\n", myvariable3)
}

//DifferentType  Demonstrates how multiple variables having different data type can be declared on a single line
func DifferentType() {
	utility.Printh(sublevel, differentTypeConstLbl)
	tabs := utility.Tabs(sublevel)
	// Multiple variables of different types
	// are declared and initialized in the single line
	var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

	// Display the value and
	// type of the variables
	fmt.Printf(tabs+"myvariable1 %d\n", myvariable1)
	fmt.Printf(tabs+"myvariable2 %s\n", myvariable2)
	fmt.Printf(tabs+"myvariable3 %f\n", myvariable3)

}

//Function  Initialize variables from the function return type
func Function() {
	utility.Printh(sublevel, functionConstLbl)
	tabs := utility.Tabs(sublevel)
	// Println returns Multiple return values
	// First return variables return bytes used
	// Second return Error occured wile processessing
	var bytesUsed, errorsGenerated = fmt.Println(tabs + "Hello World! I love Go Programming!")

	utility.Println(sublevel, "Bytes Used : ", bytesUsed)
	utility.Println(sublevel, "Error Code : ", errorsGenerated)
}

//Blank  Initialize variables from the function return type and assign to _ if the returned value is not to be used further in the code
func Blank() {
	// Println returns Multiple return values
	// First return variables return bytes used
	// Second return Error occured wile processessing
	utility.Printh(sublevel, blankConstLbl)
	tabs := utility.Tabs(sublevel)
	var bytesUsed, _ = fmt.Println(tabs + "Hello World! I love Go Programming!")
	// As the second argument need not to be used the code further , _ notify the same to the compiler
	utility.Println(sublevel, "Bytes Used : ", bytesUsed)

}
