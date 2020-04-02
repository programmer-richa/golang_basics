package exercise

import "fmt"

//SwapTwoNumbersWithoutUsingThirtVariable : Swapping simply by x,y=y,X Operation
func SwapTwoNumbersWithoutUsingThirtVariable() {
	x, y := 10, 20
	fmt.Println("X:", x, "y:", y)
	x, y = y, x
	fmt.Println("After swapping")
	fmt.Println("X:", x, "y:", y)
}
