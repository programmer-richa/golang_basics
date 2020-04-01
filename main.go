package main

import (
	"github.com/programmer-richa/golang_basics/helloworld"
	"github.com/programmer-richa/golang_basics/variables"
)

func main() {
	printSeperator("Hello World Example")
	helloworld.Greeting()
	printSeperator("Variable Declaration Ways")
	variables.VariableDeclarationWays()
	printSeperator("Variable Type Tracking")
	variables.VariableTypeTracking()
	printSeperator("Zero Values Of Basic Data Types")
	variables.ZeroValuesOfBasicDataTypes()
	printSeperator("Declaring Multiple Variables With Same Data Type In a Single Line")
	variables.DeclaringMultipleVariablesWithSameDataTypeInSingleLine()
	printSeperator("Declaring Multiple Variables With Different Data Type In a Single Line")
	variables.DeclaringMultipleVariablesWithDifferentDataTypeInSingleLine()
	printSeperator("Initializing Variables From a Function Returning Multiple Values")
	variables.InitializingVariablesFromAFunctionReturningMultipleValues()
	printSeperator("Using Blank Identifier to Overcome The Compultion to Use Variable In Code")
	variables.UsingBlankIdentifierToOvercomeTheCompultionToUseVariableInCode()
}

func printSeperator(heading string) {
	println("")
	println("")
	println(heading)
	println("__________________________________________________________________________")
	println("")
}
