package exercise

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

// Excerise Modules Listing Constants
const (
	_ = iota
	swapConst
	arraysumConst
	sliceConst
	exitConst
	swapConstLbl     = "Swapping Two Numbers"
	arraysumConstLbl = "Add Elements of 2D Array"
	sliceConstLbl    = "Demonstrates working of slice according to number of students,subjects specified by user"
	exitConstLbl     = "Exit"
)

//Block Works according to choice of user
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, swapConst, swapConstLbl)
		utility.Println(level, arraysumConst, arraysumConstLbl)
		utility.Println(level, sliceConst, sliceConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case swapConst:
			utility.Printh(sublevel, swapConstLbl)
			swap()
		case arraysumConst:
			utility.Printh(sublevel, arraysumConstLbl)
			a := [][]int{
				{1, 2, 3},
				{4, 5},
			}
			sum := add2d(&a)
			utility.Println(sublevel, "Sum = ", sum)
		case sliceConst:
			utility.Printh(level, sliceConstLbl)
			dynamicMarks()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}
