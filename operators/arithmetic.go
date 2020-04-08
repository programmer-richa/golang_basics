package operators

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

const (
	_ = iota
	additionConst
	subtractionConst
	multiplicationConst
	divisionConst
	modulusConst
	reverseConst
	additionConstLbl       = " Addition of 2 numbers"
	subtractionConstLbl    = " Subtraction of 2 numbers"
	multiplicationConstLbl = " Multiplication of 2 numbers"
	divisionConstLbl       = " Division of 2 numbers"
	modulusConstLbl        = " Modulous of 2 numbers"
	reverseConstLbl        = " Reverse Of a number"
)

//ArithmeticOperator examples of arithemetic operation
func ArithmeticOperator() {
	var choice int
	utility.Println(sublevel, "Choose From Options:")
	utility.Println(sublevel, additionConst, additionConstLbl)
	utility.Println(sublevel, subtractionConst, subtractionConstLbl)
	utility.Println(sublevel, multiplicationConst, multiplicationConstLbl)
	utility.Println(sublevel, divisionConst, divisionConstLbl)
	utility.Println(sublevel, modulusConst, modulusConstLbl)
	utility.Println(sublevel, reverseConst, reverseConstLbl)
	utility.Print(sublevel, "Enter your choice:")
	fmt.Scan(&choice)
	switch choice {
	case additionConst:
		addTwoNumbers()
	case subtractionConst:
		subtractTwoNumbers()
	case multiplicationConst:
		multiplyTwoNumbers()
	case divisionConst:
		divideTwoNumbers()
	case modulusConst:
		remainderTwoNumbers()
	case reverseConst:
		reverseOfANumber()
	}

}

func addTwoNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	c = a + b
	fmt.Printf(tabs+"Sum of %d and %d is %d", a, b, c)
}

func subtractTwoNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	c = a - b
	fmt.Printf(tabs+"Difference of %d and %d is %d", a, b, c)
}

func multiplyTwoNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	c = a * b
	fmt.Printf(tabs+"Product of %d and %d is %d", a, b, c)
}

func divideTwoNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	c = a / b
	fmt.Printf(tabs+"Division of %d and %d is %d", a, b, c)
}

func remainderTwoNumbers() {
	var a, b, c int
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter first number :")
	fmt.Scan(&a)
	utility.Print(sublevel1, "Enter second number :")
	fmt.Scan(&b)
	c = a % b
	fmt.Printf(tabs+"Remainder of %d and %d is %d", a, b, c)
}

func reverseOfANumber() {
	var orginalNumber, number, reverse, remainder int = 0, 0, 0, 0
	tabs := utility.Tabs(sublevel1)
	utility.Print(sublevel1, "Enter number :")
	fmt.Scan(&orginalNumber)
	number = orginalNumber
	for number > 0 {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number = number / 10
	}
	fmt.Printf(tabs+"Reverse of %d is %d\n", orginalNumber, number)
	fmt.Printf(tabs+"Is Palindrome : %t\n", orginalNumber == number)
}

//Addition  returns sum of values
func Addition(a ...int) int {
	utility.Println(sublevel1, a)
	utility.Println(sublevel1)
	var sum int = 0
	for _, value := range a {
		sum += value
	}
	return sum
}
