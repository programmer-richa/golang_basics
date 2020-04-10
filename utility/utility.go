// Package utility   contains common functions used in this module.
package utility

import (
	"fmt"
	"strings"
)

//Tabs   generates '\t' strings as per user requirement.
// It accepts one argument: 1. level that specifies number of tab keys to ge generated.
func Tabs(level int) string {
	if level < 1 {
		level = 0
	}
	// call built-in string function Repeat to repeat the string
	tabSpaces := strings.Repeat("\t", level)
	return tabSpaces
}

//Printh   prints heading in a standardized format.
// It accepts two arguments :
// 1. level that specifies number of tab keys to ge generated
// 2. a that holds variadic arguments to be printed
func Printh(level int, a ...interface{}) {
	tabSpaces := Tabs(level)
	arguments := tabSpaces + "\n" + tabSpaces

	for _, value := range a {
		// concatinating arguments to string after getting the value from the interface
		arguments += fmt.Sprintf("%v ", value)
	}
	arguments += tabSpaces + "\n" + tabSpaces + "__________________________________________________________________________"
	fmt.Println(arguments)

}

//Println   prints data in a standardized format and generates a newline after printing.
// It accepts two arguments :
// 1. level that specifies number of tab keys to ge generated
// 2. a that holds variadic arguments to be printed
func Println(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string after getting the value from the interface
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Println(arguments)
}

//Print   prints data in a standardized format and does not generate a newline after printing.
// It accepts two arguments :
// 1. level that specifies number of tab keys to ge generated
// 2. a that holds variadic arguments to be printed
func Print(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string after getting the value from the interface
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Print(arguments)
}
