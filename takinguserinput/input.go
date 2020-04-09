package takinguserinput

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1

//Add Add 2 numbers entered by user
func Add() {
	utility.Printh(level, "Adding numbers via user input")
	var a, b int
	utility.Print(level, "Enter First number: ")
	fmt.Scan(&a)
	utility.Print(level, "Enter second number: ")
	fmt.Scan(&b)
	sum := a + b
	utility.Println(level, "Sum of", a, "and", b, ":", sum)
}
