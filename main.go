/*
Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code.
 In Go, there are two types of program availbl one is executable and another one is the library.
  The executable programs are those programs that we can run directly from the terminal and Libraries are the package of programs
  that we can reuse them in our program. Here, the package main tells the compiler that the package must compile
  in the executable program rather than a shared library. It is the starting point of the program and also contains
   the main function in it.
*/
package main

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/arrays"
	"github.com/programmer-richa/golang_basics/constants"
	"github.com/programmer-richa/golang_basics/controlflowstatements"
	"github.com/programmer-richa/golang_basics/exercise"
	"github.com/programmer-richa/golang_basics/functions"
	"github.com/programmer-richa/golang_basics/helloworld"
	"github.com/programmer-richa/golang_basics/identifiers"
	"github.com/programmer-richa/golang_basics/keywords"
	"github.com/programmer-richa/golang_basics/operators"
	"github.com/programmer-richa/golang_basics/pointers"
	"github.com/programmer-richa/golang_basics/runtimeexample"
	"github.com/programmer-richa/golang_basics/strings"
	"github.com/programmer-richa/golang_basics/takinguserinput"
	"github.com/programmer-richa/golang_basics/utility"
	"github.com/programmer-richa/golang_basics/variables"
)

// Package Level Constants For Module Choice
const (
	// value of iota is incremented for all constants declared
	_                  = iota // iota value starts from 0 in every separate const block
	helloConst                // or helloConst= iota // Automatically helloConst will be assigned current value of iota i.e. 1
	variableConst      = iota //iota 2
	identifierConst    = iota
	keywordsConst      = iota
	constantsFlowConst = iota
	controlFlowConst   = iota
	operatorsConst     = iota
	arrayConst         = iota
	pointerConst       = iota
	userInpuConst      = iota
	functionConst      = iota
	runtimeConst       = iota
	stringConst        = iota
	exerciseConst      = iota
	exitConst
	helloConstlbl       = "Hello World"
	variableConstlbl    = "Variable"
	identifierConstlbl  = "Identifier"
	keywordsConstlbl    = "Keywords"
	constantConstlbl    = "Constants"
	controlFlowConstlbl = "Control Flow"
	operatorsConstlbl   = "Operators"
	arrayConstlbl       = "Array"
	pointerConstlbl     = "Pointer"
	userInpuConstlbl    = "Taking User Input"
	functionConstlbl    = "Functions"
	runtimeConstlbl     = "Using runtime package"
	stringConstlbl      = "String"
	exerciseConstlbl    = "Exercises"
	exitConstLbl        = "Exit"
)

// Module Level Flag
const level = 0

//main This is the entry point of the Program.
func main() {
	var choice int // local variable
	for choice != exitConst {
		//Module Level Option Display
		utility.Println(level, "Choose From Options below :")
		utility.Println(level, helloConst, helloConstlbl)
		utility.Println(level, variableConst, variableConstlbl)
		utility.Println(level, identifierConst, identifierConstlbl)
		utility.Println(level, keywordsConst, keywordsConstlbl)
		utility.Println(level, constantsFlowConst, constantConstlbl)
		utility.Println(level, controlFlowConst, controlFlowConstlbl)
		utility.Println(level, operatorsConst, operatorsConstlbl)
		utility.Println(level, arrayConst, arrayConstlbl)
		utility.Println(level, pointerConst, pointerConstlbl)
		utility.Println(level, userInpuConst, userInpuConstlbl)
		utility.Println(level, functionConst, functionConstlbl)
		utility.Println(level, runtimeConst, runtimeConstlbl)
		utility.Println(level, stringConst, stringConstlbl)
		utility.Println(level, exerciseConst, exerciseConstlbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")
		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice
		switch choice {
		case helloConst:
			helloworld.Greeting()
		case variableConst:
			variables.Block()
		case identifierConst:
			identifiers.Usage()
		case keywordsConst:
			keywords.Usage()
		case constantsFlowConst:
			constants.Usage()
		case controlFlowConst:
			controlflowstatements.Block()
		case operatorsConst:
			operators.Block()
		case arrayConst:
			arrays.Block()
		case pointerConst:
			pointers.Block()
		case userInpuConst:
			takinguserinput.Add()
		case functionConst:
			functions.Block()
		case runtimeConst:
			runtimeexample.Data()
		case stringConst:
			strings.ToByteArray()
		case exerciseConst:
			exercise.Block()
		case exitConst:
			utility.Println(level, "Exiting Program....")
		default:
			utility.Println(level, "Invalid choice")
		}
	}
}
