/*
As the name CONSTANTS suggests means fixed, in programming languages also it is same i.e., once the value of constant is defined it cannot be modified further. There can be any basic data types of constant like an integer constant, a floating constant, a character constant, or a string literal.

How to declare?
Constant are declared like variables but in using a const keyword as a prefix to declare constant with a specific type. It cannot be declare using := syntax.

Untyped and Typed Numeric Constants:
Typed constants work like immutable variables can inter-operate only with the same type and untyped constants work like literals can inter-operate with similar types. Constants can be declared with or without a type in Go. Following is the example which show typed and untyped numeric constants that are both named and unnamed.

Following is list of constant in Go Language:

Numeric Constant (Integer constant, Floating constant, Complex constant)
String literals
Boolean Constant
*/
// Package constants   demonstrates working of global and local scoped constants, typed and untyped constants, and working of iota.
package constants

import "github.com/programmer-richa/golang_basics/utility"

const (
	_   = iota // iota value starts from 0 in every separate const block
	red        // or red= iota
	// Automatically red will be assigned current value of iota i.e. 1
	// because go compiler assigns same value to the next unassigned constanst as it is assigned to prev constant
	blue
)

const (
	same = 1 // Automatically duplicate will be assigned 1 i.e. 1 because go compiler assigns same value to the next unassigned constanst as it is assigned to prev constant
	duplicate
	other = iota // iota is incremented for every constant declared in a block, so value of other will be 2
)

// Module Level Flag is used for indentation Purpose
const level = 1

// PI value of PI
const PI = 3.14 //Untyped

//Usage   demonstrates the scope and type of Constants
func Usage() {
	utility.Printh(level, "Accessing Package and Global Scope Constants")
	const AUTHOR string = "Richa" //typed
	utility.Println(level, "Local Scope AUTHOR", AUTHOR)

	utility.Println(level, "Package Scope PT", PI)

	utility.Println(level, "Red :", red)
	utility.Println(level, "blue :", blue)
	utility.Println(level, "same :", same)
	utility.Println(level, "duplicate :", duplicate)
	utility.Println(level, "other :", other)
}
