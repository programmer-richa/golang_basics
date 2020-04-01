package variables

import (
	"fmt"
)

//VariableDeclarationWays : It demonstrates alternatives to declare and initialize variables
func VariableDeclarationWays() {
	// declared variables must be used at least once in the program, otherwise compiler generated CTE
	// Way 1: Declare and initialize variables explicitly
	var i int // declaration of int type variable i
	i = 10    // assigning value to int
	fmt.Println("Integer value assigned to variable i = ", i)

	// Way 2: Declare and initialize variables explicitly on the same line
	var f float32 = 25.46 // declaration of float type variable f for 32 bit memory
	fmt.Println("Float (32- bit )value assigned to variable f = ", f)

	var f1 float64 // declaration of float type variable f for 64 bit memory
	f1 = 25.46
	fmt.Println("Float (64- bit )value assigned to variable f = ", f1)

	//Way 3: Implicit variable declaration, here compiler will interpret the data type based on the value assigned to the variable
	// This is most frequent way of declaring variables
	firstName := "Richa"
	fmt.Println("First Name : " + firstName)

	var b bool = true // Boolean Data Type
	fmt.Println("Boolean value : ", b)

	var realNumber float32 = 5
	var imaginaryNumber float32 = 4
	// complex function expects arguments to be float
	complexNumber := complex(realNumber, imaginaryNumber)
	fmt.Println("Complex Number : ", complexNumber)

	// retrieving real and imaginary part of a complex number
	fmt.Println("Real value  : ", real(complexNumber))
	fmt.Println("Imaginary value  : ", imag(complexNumber))
}

//VariableTypeTracking : It demonstrates value and type tracking of a variable
func VariableTypeTracking() {
	// Variable declared and
	// initialized without the
	// explicit type
	var myvariable1 = 20
	var myvariable2 = "GeeksforGeeks"
	var myvariable3 = 34.80

	// Display the value and the
	// type of the variables
	fmt.Printf("The value of myvariable1 is : %d\n", myvariable1)

	fmt.Printf("The type of myvariable1 is : %T\n", myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n", myvariable2)

	fmt.Printf("The type of myvariable2 is : %T\n", myvariable2)

	fmt.Printf("The value of myvariable3 is : %f\n", myvariable3)

	fmt.Printf("The type of myvariable3 is : %T\n", myvariable3)
}

//ZeroValuesOfBasicDataTypes :  Displays default values of basic data types
func ZeroValuesOfBasicDataTypes() {
	// Variable declared and
	// initialized without expression
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
	// fmt.Printf("The value of %T is : %d\n", int1, int1)
	printDataTypeAndValue(int1, "%d")
	printDataTypeAndValue(int2, "%d")
	printDataTypeAndValue(int3, "%d")
	printDataTypeAndValue(int4, "%d")
	printDataTypeAndValue(int5, "%d")
	printDataTypeAndValue(uint1, "%d")
	printDataTypeAndValue(uint2, "%d")
	printDataTypeAndValue(uint3, "%d")
	printDataTypeAndValue(uint4, "%d")
	printDataTypeAndValue(uint5, "%d")
	printDataTypeAndValue(float1, "%f")
	printDataTypeAndValue(float2, "%f")
	printDataTypeAndValue(string1, "%s")
	printDataTypeAndValue(bool1, "%t")
	printDataTypeAndValue(complex1, "%g")
	printDataTypeAndValue(complex2, "%g")
}

// here first argument  a interface{} suggests that a variable of any data type can be passed
// second argument stringFmt string restricts the data type of the argument to be string
func printDataTypeAndValue(a interface{}, stringFmt string) {
	fmt.Printf("The value of %T is : "+stringFmt+"\n", a, a)
}

//DeclaringMultipleVariablesWithSameDataTypeInSingleLine :  Demonstrates how multiple variables having same data type can be declared on a single line
func DeclaringMultipleVariablesWithSameDataTypeInSingleLine() {
	// Multiple variables of the same type
	// are declared and initialized
	// in the single line
	var myvariable1, myvariable2, myvariable3 int = 2, 454, 67

	// Display the values of the variables
	printDataTypeAndValue(myvariable1, "%d")
	printDataTypeAndValue(myvariable2, "%d")
	printDataTypeAndValue(myvariable3, "%d")
}

//DeclaringMultipleVariablesWithDifferentDataTypeInSingleLine :  Demonstrates how multiple variables having different data type can be declared on a single line
func DeclaringMultipleVariablesWithDifferentDataTypeInSingleLine() {
	// Multiple variables of different types
	// are declared and initialized in the single line
	var myvariable1, myvariable2, myvariable3 = 2, "GFG", 67.56

	// Display the value and
	// type of the variables
	printDataTypeAndValue(myvariable1, "%d")
	printDataTypeAndValue(myvariable2, "%s")
	printDataTypeAndValue(myvariable3, "%f")

}

//InitializingVariablesFromAFunctionReturningMultipleValues :  Initialize variables from the function return type
func InitializingVariablesFromAFunctionReturningMultipleValues() {
	// Println returns Multiple return values
	// First return variables return bytes used
	// Second return Error occured wile processessing
	var bytesUsed, errorsGenerated = fmt.Println("Hello World! I love Go Programming!")

	fmt.Println("Bytes Used : ", bytesUsed)
	fmt.Println("Error Code : ", errorsGenerated)
}

//UsingBlankIdentifierToOvercomeTheCompultionToUseVariableInCode :  Initialize variables from the function return type and assign to _ if the returned value is not to be used further in the code
func UsingBlankIdentifierToOvercomeTheCompultionToUseVariableInCode() {
	// Println returns Multiple return values
	// First return variables return bytes used
	// Second return Error occured wile processessing
	var bytesUsed, _ = fmt.Println("Hello World! I love Go Programming!")
	// As the second argument need not to be used the code further , _ notify the same to the compiler
	fmt.Println("Bytes Used : ", bytesUsed)

}
