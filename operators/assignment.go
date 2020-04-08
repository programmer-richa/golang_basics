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

//AssignmentOperator examples of arithemetic operation
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
