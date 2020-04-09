package operators

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1
const sublevel1 = sublevel + 1

// Excerise Modules Listing Constants
const (
	_ = iota
	arithematicConst
	assignmentConst
	exitConst
	arithematicConstLbl = "Arithematic"
	assignmentConstLbl  = "Assignment"
	exitConstLbl        = "Exit"
)

//Block This function enables user to choose from the list of options available to test variety of operators implemented in this sub module
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, arithematicConst, arithematicConstLbl)
		utility.Println(level, assignmentConst, assignmentConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case arithematicConst:
			ArithmeticOperator()
		case assignmentConst:
			AssignmentOperator()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}
