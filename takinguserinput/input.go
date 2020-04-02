package takinguserinput

import "fmt"

// AddingTwoNumbers Add 2 numbers entered by user
func AddingTwoNumbers() {
	var a, b int
	fmt.Print("Enter First number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	sum := a + b
	fmt.Println("Sum of", a, "and", b, ":", sum)
}
