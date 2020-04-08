package utility

import (
	"fmt"
	"strings"
)

//Tabs level wise spaces
func Tabs(level int) string {
	if level < 1 {
		level = 0
	}
	tabSpaces := strings.Repeat("\t", level)
	return tabSpaces
}

//Printh  print tabs and arguments passed
func Printh(level int, a ...interface{}) {
	tabSpaces := Tabs(level)
	arguments := tabSpaces + "\n" + tabSpaces

	for _, value := range a {
		// concatinating arguments to string
		arguments += fmt.Sprintf("%v ", value)
	}
	arguments += tabSpaces + "\n" + tabSpaces + "__________________________________________________________________________"
	fmt.Println(arguments)

}

//Println  print tabs and arguments passed
func Println(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Println(arguments)
}

//Print  print tabs and arguments passed
func Print(level int, a ...interface{}) {

	arguments := Tabs(level)

	for _, value := range a {
		// concatinating arguments to string
		arguments += fmt.Sprintf("%v ", value)
	}
	fmt.Print(arguments)
}
