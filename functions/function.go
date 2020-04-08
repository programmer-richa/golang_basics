package functions

import (
	"fmt"
	"sort"
	"strings"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

// Array Modules Listing Constants
const (
	_ = iota
	maxConst
	swapConst
	swapreferenceConst
	implodeConst
	explodeConst
	arraySumConst
	exitConst
	maxConstLbl           = "Max of 2 numbers"
	swapConstLbl          = "Swap 2 numbers by value"
	swapreferenceConstLbl = "Swap 2 numbers by reference"
	implodeConstLbl       = "Join String"
	explodeConstLbl       = "Split String"
	arraySumConstLbl      = "Sum of 1D array Elements"
	exitConstLbl          = "Exit"
)

//Block Works according to choice of user
func Block() {
	var a, b, choice = 10, 20, 0 // local variable
	tabs := utility.Tabs(sublevel)
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, maxConst, maxConstLbl)
		utility.Println(level, swapConst, swapConstLbl)
		utility.Println(level, swapreferenceConst, swapreferenceConstLbl)
		utility.Println(level, implodeConst, implodeConstLbl)
		utility.Println(level, explodeConst, explodeConstLbl)
		utility.Println(level, arraySumConst, arraySumConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case maxConst:
			utility.Printh(sublevel, maxConstLbl)
			fmt.Printf(tabs+"Max of %d and %d is %d\n", a, b, Max(a, b))
		case swapConst:
			utility.Printh(sublevel, swapConstLbl)
			utility.Println(sublevel, "In main before swapping x = ", a)
			utility.Println(sublevel, "In main before swapping y = ", b)
			SwapByValue(a, b)
			utility.Println(sublevel, "In main after swapping x = ", a)
			utility.Println(sublevel, "In main after swapping y = ", b)
		case swapreferenceConst:
			utility.Printh(sublevel, swapreferenceConstLbl)
			utility.Println(sublevel, "In main before swapping x = ", a)
			utility.Println(sublevel, "In main before swapping y = ", b)
			SwapByReference(&a, &b)
			utility.Println(sublevel, "In main after swapping x = ", a)
			utility.Println(sublevel, "In main after swapping y = ", b)
		case implodeConst:
			utility.Printh(sublevel, implodeConstLbl)
			line := Implode(",", "I", "love", "go", "language", ".")
			utility.Println(sublevel, "Line generated : ", line)
		case explodeConst:
			line := "I,love,go,language,."
			utility.Printh(sublevel, explodeConstLbl)
			tokens := Explode(",", line)
			utility.Println(sublevel, "Tokens generated : ", tokens)
			utility.Println(sublevel, "Tokens length : ", len(tokens))
		case arraySumConst:
			utility.Printh(sublevel, arraySumConstLbl)
			arr := []int{10, 20, 30, 40, 50}
			total := SumOf1DArray(arr)
			utility.Println(sublevel, "Total of arr : ", arr, " is ", total)
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Explode : Variadic Functions that splits a string with given separator and returns string array
func Explode(separator string, str string) []string {
	return strings.Split(str, separator)
}

//Implode : Variadic Functions that joins a string with given separator and returns string array
func Implode(separator string, str ...string) string {
	return strings.Join(str, separator)
}

//SwapByValue : Pass By value , separate copy of arguments is created in the memory,so the changes made in the formal arguments are not reflected in the actual arguments
func SwapByValue(x int, y int) {
	temp := x
	x = y
	y = temp
	utility.Println(sublevel, "In swapByValue x = ", x)
	utility.Println(sublevel, "In swapByValue y = ", y)
}

//SwapByReference : Pass By Reference , address of arguments passed to the function, so the changes made in the formal arguments are reflected in the actual arguments
func SwapByReference(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
	utility.Println(sublevel, "In swapByValue x = ", *x)
	utility.Println(sublevel, "In swapByValue y = ", *y)
}

// Max : returns max of two numbers
func Max(x int, y int) int {
	z := y
	if x > y {
		z = x
	}
	return z
}

// SumOf1DArray : returns sum of of 1D Array
func SumOf1DArray(a []int) int {
	sizeOnArray := len(a)
	total := 0
	for i := 0; i < sizeOnArray; i++ {
		total += a[i]
	}

	return total
}

// Sort1DArrayByValue : sorts array
func Sort1DArrayByValue(a []int) {
	sort.Ints(a)

}
