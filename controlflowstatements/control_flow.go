// Package controlflowstatements implements various control flow structures available in go lang.
package controlflowstatements

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag is used for indentation purpose.
const level = 1
const sublevel = level + 1

// Control Flow Statements Modules Listing Constants.
const (
	_ = iota
	comparisionConst
	optionConst
	noOptionConst
	expressionConst
	typeConst
	fallthroughConst
	defaultConst
	loopConst
	breakConst
	continueConst
	lableConst
	gOtoConst
	infiniteConst
	whileConst
	rangeConst
	onStringConst
	mapConst
	channelConst
	exitConst
	comparisionConstLbl = "If ElseIf Else Example"
	optionConstLbl      = " Switch With OptionalStatement Example"
	noOptionConstLbl    = "Switch Without Optional Statement Example"
	expressionConstLbl  = "Switch Expression Example"
	typeConstLbl        = "Switch Type Variable Example"
	fallthroughConstLbl = "Switch Case with fallthrough"
	defaultConstLbl     = "Switch Without Option variable has an default value true"
	loopConstLbl        = "For Loop Example"
	breakConstLbl       = "For Loop With Break Example"
	continueConstLbl    = "For Loop With Continue Example"
	lableConstLbl       = "For Loop With Labled Continue Example"
	gOtoConstLbl        = "For Loop With GoTo Example"
	infiniteConstLbl    = "For Infinite Loop Example"
	whileConstLbl       = "For Loop Like While Example"
	rangeConstLbl       = "For Loop On Range Example"
	onStringConstLbl    = "For Loop On String Example"
	mapConstLbl         = "For Loop On Map Example"
	channelConstLbl     = "For Loop on Channel Example"
	exitConstLbl        = "Exit"
)

//Block   enables user to choose from the list of options available to test variety of control flow statements implemented
// in this sub module
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, comparisionConst, comparisionConstLbl)
		utility.Println(level, optionConst, optionConstLbl)
		utility.Println(level, noOptionConst, noOptionConstLbl)
		utility.Println(level, expressionConst, expressionConstLbl)
		utility.Println(level, typeConst, typeConstLbl)
		utility.Println(level, fallthroughConst, fallthroughConstLbl)
		utility.Println(level, defaultConst, defaultConstLbl)
		utility.Println(level, loopConst, loopConstLbl)
		utility.Println(level, breakConst, breakConstLbl)
		utility.Println(level, continueConst, continueConstLbl)
		utility.Println(level, lableConst, lableConstLbl)
		utility.Println(level, gOtoConst, gOtoConstLbl)
		utility.Println(level, infiniteConst, infiniteConstLbl)
		utility.Println(level, whileConst, whileConstLbl)
		utility.Println(level, rangeConst, rangeConstLbl)
		utility.Println(level, onStringConst, onStringConstLbl)
		utility.Println(level, mapConst, mapConstLbl)
		utility.Println(level, channelConst, channelConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case comparisionConst:
			Comparision()
		case optionConst:
			Options()
		case noOptionConst:
			NoOptions()
		case expressionConst:
			Expression()
		case typeConst:
			Type()
		case fallthroughConst:
			Fallthrough()
		case defaultConst:
			DefaultChoice()
		case loopConst:
			Loop()
		case breakConst:
			Break()
		case continueConst:
			Continue()
		case lableConst:
			Lable()
		case gOtoConst:
			GOto()
		case infiniteConst:
			Infinite()
		case whileConst:
			While()
		case rangeConst:
			Range()
		case onStringConst:
			OnString()
		case mapConst:
			Map()
		case channelConst:
			Channel()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

// Comparision   uses if-else if - else ladder to compare 2 int variables.
func Comparision() {
	utility.Printh(sublevel, comparisionConstLbl)
	a, b := 10, 20 // This is recommended way of declaring local variables
	if a > b {
		utility.Println(sublevel, a, "is greater than", b)
	} else if a < b {
		utility.Println(sublevel, a, "is smaller than", b)
	} else {
		utility.Println(sublevel, "Both are equal")
	}
}

// Options  demonstrates the working of switch block that contains the optional statement
// that inititializes the choice variable for example "day" and expression to test the truthiness for case evaluation.
func Options() {
	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	utility.Printh(level, optionConstLbl)
	switch day := 4; day {
	case 1:
		utility.Println(sublevel, "Monday")
	case 2:
		utility.Println(sublevel, "Tuesday")
	case 3:
		utility.Println(sublevel, "Wednesday")
	case 4:
		utility.Println(sublevel, "Thursday")
	case 5:
		utility.Println(sublevel, "Friday")
	case 6:
		utility.Println(sublevel, "Saturday")
	case 7:
		utility.Println(level, "Sunday")
	default:
		utility.Println(sublevel, "Invalid")
	}
}

// Fallthrough   demonstrates the working of switch block that specifies fallthrough statements,
// and contains  the optional statement that inititializes the choice variable for example "day"
// and expression to test the truthiness for case evaluation.
func Fallthrough() {
	utility.Printh(sublevel, fallthroughConstLbl)
	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	switch day := 4; day {
	case 1:
		utility.Println(sublevel, "Monday")
	case 2:
		utility.Println(sublevel, "Tuesday")
	case 3:
		utility.Println(sublevel, "Wednesday")
	case 4:
		utility.Println(sublevel, "Thursday")
		fallthrough
	case 5:
		utility.Println(sublevel, "Friday")
	case 6:
		utility.Println(sublevel, "Saturday")
	case 7:
		utility.Println(sublevel, "Sunday")
	default:
		utility.Println(sublevel, "Invalid")
	}
}

// NoOptions  demonstrates the working of switch block that does not contain  the optional statement
// that inititializes the choice variable  and expression to test the truthiness for case evaluation.
func NoOptions() {
	utility.Printh(sublevel, noOptionConstLbl)
	var value int = 2

	// Switch statement without an
	// optional statement and
	// expression
	switch {
	case value == 1:
		utility.Println(sublevel, "Hello")
	case value == 2:
		utility.Println(sublevel, "Bonjour")
	case value == 3:
		utility.Println(sublevel, "Namstay")
	default:
		utility.Println(sublevel, "Invalid")
	}
}

// DefaultChoice   demonstrates the default choice "True" of switch block that does not contain  the optional statement
// that inititializes the choice variable  and expression to test the truthiness for case evaluation.
func DefaultChoice() {
	utility.Printh(sublevel, defaultConstLbl)
	switch {
	case true:
		utility.Println(sublevel, "True block")
	case false:
		utility.Println(sublevel, "false block")
	default:
		utility.Println(sublevel, "Invalid")
	}
}

// Expression  demonstrates the abitity of case statement to accept multiple expressions.
func Expression() {
	utility.Printh(sublevel, expressionConstLbl)
	var value string = "three"

	// Switch statement without default statement
	// Multiple values in case statement
	switch value {
	case "one":
		utility.Println(sublevel, "C#")
	case "two", "three":
		utility.Println(sublevel, "Go")
	case "four", "five", "six":
		utility.Println(sublevel, "Java")
	}
}

// Type  demonstrates the abitity of switch statement to execute according to the data type of the optional statement variable.
func Type() {
	tabs := utility.Tabs(level + 1)
	utility.Printh(sublevel, typeConstLbl)
	var value interface{}
	switch q := true; value.(type) {
	case bool:
		utility.Println(sublevel, "value is of boolean type")
	case float64:
		utility.Println(sublevel, "value is of float64 type")
	case int:
		utility.Println(sublevel, "value is of int type")
	default:
		fmt.Printf(tabs+"value is of type: %T", q)

	}
}

// Loop   implements simple for loop to print values from 1 to 10.
func Loop() {
	utility.Printh(sublevel, loopConstLbl)
	for i := 1; i < 11; i++ {
		utility.Println(sublevel, i)
	}
}

// Break  implements break statement to exit for loop when a specific condition is met.
func Break() {
	utility.Printh(sublevel, breakConstLbl)
	for i := 1; i < 11; i++ {
		if i == 5 {
			break
		}
		utility.Println(sublevel, i)
	}
}

// Continue  implements continue statement to skip for loop iteration when a specific condition is met.
func Continue() {
	utility.Printh(sublevel, continueConstLbl)
	for i := 1; i < 11; i++ {
		if i == 5 {
			continue
		}
		utility.Println(sublevel, i)
	}
}

// Lable  implements labeled continue statement to skip for loop iteration
// and continue with the labeled statement when a specific condition is met.
func Lable() {
	utility.Printh(sublevel, lableConstLbl)
Lable1:
	for i := 1; i < 5; i++ {
		utility.Println(sublevel, "Value of i: ", i)
		for j := 1; j <= 5; j++ {
			if j == 3 {
				continue Lable1
			}
			utility.Println(sublevel, "Value of j: ", j)
		}

	}
}

// GOto  implements goto statement to jump to the specified labeled position.
func GOto() {
	tabs := utility.Tabs(level + 1)
	utility.Printh(sublevel, gOtoConstLbl)
	var x int = 0

	// for loop work as a while loop
Lable1:
	for x < 8 {
		if x == 5 {

			// using goto statement
			x++
			goto Lable1
		}
		fmt.Printf(tabs+"value is: %d\n", x)
		x++
	}
}

// Infinite   demonstrates the working of infinite for loop.
func Infinite() {
	utility.Printh(sublevel, infiniteConstLbl)
	i := 1
	for {
		utility.Println(sublevel, i)
		i++
		if i == 5 {
			break
		}
	}
}

// While   demonstrates the working of for loop which contains condition block only
// and omits initialization and conditional block, i.e. similar to while loop.
func While() {
	utility.Printh(sublevel, whileConstLbl)
	i := 1
	for i < 5 {
		utility.Println(sublevel, i)
		i++
	}
}

// Range   demonstrates how for loop can be used to iterate over range of values in a slice.
func Range() {
	utility.Printh(sublevel, rangeConstLbl)
	// Here rvariable is a array
	rvariable := []string{"Ram", "Sham", "Month"}

	// key and value stores the value of rvariable
	// key store index number of individual string and
	// value store individual string of the given array
	for key, value := range rvariable {
		utility.Println(sublevel, key, value)
	}
}

// OnString   demonstrates how for loop can be used to iterate over range of values in a string.
func OnString() {
	tabs := utility.Tabs(sublevel)
	utility.Printh(sublevel, onStringConstLbl)
	// String as a range in the for loop
	for i, j := range "Richa" {
		fmt.Printf(tabs+"The index number of %c is %d\n", j, i)
	}
}

// Map   demonstrates how for loop can be used to iterate over range of values in a map.
func Map() {
	utility.Printh(sublevel, mapConstLbl)
	// using maps
	mmap := map[int]string{
		22: "Ram",
		33: "Mohan",
		44: "Sohan",
	}
	for key, value := range mmap {
		utility.Println(sublevel, key, value)
	}
}

// Channel   demonstrates how for loop can be used to iterate over range of values in a channel.
func Channel() {
	utility.Printh(sublevel, channelConstLbl)
	// using channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		utility.Println(sublevel, i)
	}
}
