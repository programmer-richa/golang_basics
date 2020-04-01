package pointers

import "fmt"

//DefaultValueOfPointer : default value of pointer type is <nil>
func DefaultValueOfPointer() {
	// Declaring Pointer
	var firstName *string

	fmt.Println("Default Value of : ", firstName)
}

//PointerDeclarationAndInitialization : declare and initialize pointer variable
func PointerDeclarationAndInitialization() {
	// Declaring Pointer
	var firstName *string
	// Initializing Pointer
	firstName = new(string)
	// Assigning Value to Pointer
	*firstName = "Richa"
	fmt.Println("Value stored at the address refered by pointer : ", *firstName)
	fmt.Println("Address of pointer : ", &firstName)
	fmt.Println("Value stored in pointer : ", firstName)
}

//PointerDeclarationAndInitializationWithVariableReference : declare and initialize pointer variable
func PointerDeclarationAndInitializationWithVariableReference() {
	// Declaring Pointer
	var firstName string = "Richa"
	// Initializing Pointer
	// var firstNamePtr *string = &firstName
	// or
	firstNamePtr := &firstName
	fmt.Println("Value stored in variable refered by pointer : ", firstName)
	fmt.Println("Address of variable refered by pointer : ", &firstName)

	fmt.Println("Value stored at the address refered by pointer : ", *firstNamePtr)
	fmt.Println("Address of pointer : ", &firstNamePtr)
	fmt.Println("Value stored in pointer : ", firstNamePtr)
}
