package operators

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

const (
	_ = iota
	additionAssignmentConst
	subtractionAssignmentConst
	multiplicationAssignmentConst
	divisionAssignmentConst
	modulusAssignmentConst

	additionAssignmentConstLbl       = "Addition Assignment"
	subtractionAssignmentConstLbl    = "Subtraction Assignment"
	multiplicationAssignmentConstLbl = "Multiplication Assignment"
	divisionAssignmentConstLbl       = "Division Assignment"
	modulusAssignmentConstLbl        = "Modulous Assignment"
)

//AssignmentOperator This function enables user to choose from the list of options available to test variety of assignment operators implemented in this sub module
func AssignmentOperator() {
	var choice int
	utility.Println(sublevel, "Choose option:")
	utility.Println(sublevel, additionAssignmentConst, additionAssignmentConstLbl)
	utility.Println(sublevel, subtractionAssignmentConst, subtractionAssignmentConstLbl)
	utility.Println(sublevel, multiplicationAssignmentConst, multiplicationAssignmentConstLbl)
	utility.Println(sublevel, divisionAssignmentConst, divisionAssignmentConstLbl)
	utility.Println(sublevel, modulusAssignmentConst, modulusAssignmentConstLbl)
	utility.Print(sublevel, "Enter your choice:")
	fmt.Scan(&choice)
	switch choice {
	case additionAssignmentConst:
		addAssignmentNumbers()
	case subtractionAssignmentConst:
		subtractAssignmentNumbers()
	case multiplicationAssignmentConst:
		multiplyAssignmentNumbers()
	case divisionAssignmentConst:
		divideAssignmentNumbers()
	case modulusAssignmentConst:
		remainderAssignmentNumbers()

	}
}

//addAssignmentNumbers This package scoped function accepts 2 integeral values from user and assigns the sum of entered values to the first variable
func addAssignmentNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	c = a
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	a += b
	fmt.Printf(tabs+"Before Addition a = %d\t After Addition a= %d\n", c, a)
}

//subtractAssignmentNumbers This package scoped function accepts 2 integeral values from user and assigns the difference of entered values to the first variable
func subtractAssignmentNumbers() {
	tabs := utility.Tabs(sublevel1)
	var a, b, c int
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	c = a
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	a -= b
	fmt.Printf(tabs+"Before Subtraction a = %d\t After Subtraction a= %d\n", c, a)
}

//multiplyAssignmentNumbers This package scoped function accepts 2 integeral values from user and assigns the product of entered values to the first variable
func multiplyAssignmentNumbers() {
	tabs := utility.Tabs(sublevel1)
	var a, b, c int
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	c = a
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	a -= b
	fmt.Printf(tabs+"Before Multiplication a = %d\t After Multiplication a= %d\n", c, a)
}

//divideAssignmentNumbers This package scoped function accepts 2 integeral values from user and assigns the quotent of entered values to the first variable
func divideAssignmentNumbers() {
	tabs := utility.Tabs(sublevel1)
	var a, b, c int
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	c = a
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	a -= b
	fmt.Printf(tabs+"Before Division a = %d\t After Division a= %d\n", c, a)
}

//remainderAssignmentNumbers This package scoped function accepts 2 integeral values from user and assigns the remainder of entered values to the first variable
func remainderAssignmentNumbers() {
	tabs := utility.Tabs(sublevel1)
	var a, b, c int
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	c = a
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	a -= b
	fmt.Printf(tabs+"Before Modulus a = %d\t After Modulus a= %d\n", c, a)
}
