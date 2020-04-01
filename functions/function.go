package functions

import (
	"fmt"
	"sort"
	"strings"
)

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
	fmt.Println("In swapByValue x = ", x)
	fmt.Println("In swapByValue y = ", y)
}

//SwapByReference : Pass By Reference , address of arguments passed to the function, so the changes made in the formal arguments are reflected in the actual arguments
func SwapByReference(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
	fmt.Println("In swapByValue x = ", *x)
	fmt.Println("In swapByValue y = ", *y)
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
