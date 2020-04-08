package controlflowstatements

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

//Control Flow Statements Modules Listing Constants
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

//Block Works according to choice of user
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

// Comparision : If elseif statement
func Comparision() {
	utility.Printh(sublevel, comparisionConstLbl)
	var a int = 10
	var b int = 20
	if a > b {
		utility.Println(sublevel, a, "is greater than", b)
	} else if a < b {
		utility.Println(sublevel, a, "is smaller than", b)
	} else {
		utility.Println(sublevel, "Both are equal")
	}
}

// Options : switch statement
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

// Fallthrough : switch statement allows fallthrough when specified explicitly
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

// NoOptions : switch without choice variable
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

// DefaultChoice : switch without choice variable
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

// Expression : switch with multiple values in case statement
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

// Type : switch based on data type of variable
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

// Loop : Simple Loop
func Loop() {
	utility.Printh(sublevel, loopConstLbl)
	for i := 1; i < 11; i++ {
		utility.Println(sublevel, i)
	}
}

// Break : break statement example
func Break() {
	utility.Printh(sublevel, breakConstLbl)
	for i := 1; i < 11; i++ {
		if i == 5 {
			break
		}
		utility.Println(sublevel, i)
	}
}

// Continue : continue statement example
func Continue() {
	utility.Printh(sublevel, continueConstLbl)
	for i := 1; i < 11; i++ {
		if i == 5 {
			continue
		}
		utility.Println(sublevel, i)
	}
}

// Lable : continue statement example
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

// GOto : goto statement example
func GOto() {
	tabs := utility.Tabs(level + 1)
	utility.Printh(sublevel, gOtoConstLbl)
	var x int = 0

	// for loop work as a while loop
Lable1:
	for x < 8 {
		if x == 5 {

			// using goto statement
			x = x + 1
			goto Lable1
		}
		fmt.Printf(tabs+"value is: %d\n", x)
		x++
	}
}

// Infinite : Infinite Loop
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

// While : For loop with only condition statement
func While() {
	utility.Printh(sublevel, whileConstLbl)
	i := 1
	for i < 5 {
		utility.Println(sublevel, i)
		i++
	}
}

// Range : Loops over a array or range of values
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

// OnString : Loops over a string
func OnString() {
	tabs := utility.Tabs(sublevel)
	utility.Printh(sublevel, onStringConstLbl)
	// String as a range in the for loop
	for i, j := range "Richa" {
		fmt.Printf(tabs+"The index number of %c is %d\n", j, i)
	}
}

// Map : Loops over a map
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

// Channel : Loops over channel values
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
