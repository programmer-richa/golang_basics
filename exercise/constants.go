//Package exercise contains miscellaneous examples of topic covered in go lang.
package exercise

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag is used for indentation Purpose
const level = 1
const sublevel = level + 1

// Excerise Modules Listing Constants
const (
	_ = iota
	swapConst
	slicesumConst
	sliceConst
	exitConst
	swapConstLbl     = "Swapping Two Numbers"
	slicesumConstLbl = "Add Elements of 2D Slice"
	sliceConstLbl    = "Demonstrates working of slice according to number of students,subjects specified by user"
	exitConstLbl     = "Exit"
)

//Block  enables user to choose from the list of options available to test variety of exercises implemented in this sub module.
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, swapConst, swapConstLbl)
		utility.Println(level, slicesumConst, slicesumConstLbl)
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
		case slicesumConst:
			utility.Printh(sublevel, slicesumConstLbl)
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
