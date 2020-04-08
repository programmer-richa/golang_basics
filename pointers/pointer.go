package pointers

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

// Array Modules Listing Constants
const (
	_ = iota
	zeroConst
	initializeConst
	referenceConst
	exitConst
	zeroConstLbl       = "Pointer With Zero Values"
	initializeConstLbl = "Declaring And Initializing Pointer"
	referenceConstLbl  = "Declaration And Initialization With Variable Reference"
	exitConstLbl       = "Exit"
)

//Block Works according to choice of user
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, zeroConst, zeroConstLbl)
		utility.Println(level, initializeConst, initializeConstLbl)
		utility.Println(level, referenceConst, referenceConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case zeroConst:
			Zero()
		case initializeConst:
			Initialize()
		case referenceConst:
			Reference()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Zero : default value of pointer type is <nil>
func Zero() {
	utility.Printh(sublevel, zeroConstLbl)
	// Declaring Pointer
	var firstName *string

	utility.Println(sublevel, " Default Value of : ", firstName)
}

//Initialize : declare and initialize pointer variable
func Initialize() {
	utility.Printh(sublevel, initializeConstLbl)
	// Declaring Pointer
	var firstName *string
	// Initializing Pointer
	firstName = new(string)
	// Assigning Value to Pointer
	*firstName = "Richa"
	utility.Println(sublevel, " Value stored at the address refered by pointer : ", *firstName)
	utility.Println(sublevel, " Address of pointer : ", &firstName)
	utility.Println(sublevel, " Value stored in pointer : ", firstName)
}

//Reference : declare and initialize pointer variable
func Reference() {
	utility.Printh(sublevel, referenceConstLbl)
	// Declaring Pointer
	var firstName string = "Richa"
	// Initializing Pointer
	// var firstNamePtr *string = &firstName
	// or
	firstNamePtr := &firstName
	utility.Println(sublevel, " Value stored in variable refered by pointer : ", firstName)
	utility.Println(sublevel, " Address of variable refered by pointer : ", &firstName)

	utility.Println(sublevel, " Value stored at the address refered by pointer : ", *firstNamePtr)
	utility.Println(sublevel, " Address of pointer : ", &firstNamePtr)
	utility.Println(sublevel, " Value stored in pointer : ", firstNamePtr)
}
