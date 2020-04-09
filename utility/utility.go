package utility

import (
	"fmt"
	"strings"
)

//Tabs   This function generates '\t' strings as per user requirement. It accepts argument: level:Specifies Number of tab keys to ge generated
func Tabs(level int) string {
	if level < 1 {
		level = 0
	}
	// call built-in string function Repeat to repeat the string
	tabSpaces := strings.Repeat("\t", level)
	return tabSpaces
}

//Printh   This function prints heading in a standardized format. It accepts arguments : level:Specifies Number of tab keys to ge generated and a: Holds Variadic arguments to be printed
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

//Println   This function prints data in a standardized format and generates a newline after printing. It accepts arguments : level:Specifies Number of tab keys to ge generated and a: Holds Variadic arguments to be printed
func Println(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string after getting the value from the interface
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Println(arguments)
}

//Print   This function prints data in a standardized format without generating a newline after command execution. It accepts arguments : level:Specifies Number of tab keys to ge generated and a: Holds Variadic arguments to be printed
func Print(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string after getting the value from the interface
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Print(arguments)
}
