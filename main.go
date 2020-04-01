package main

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/arrays"
	"github.com/programmer-richa/golang_basics/controlflowstatements"
	"github.com/programmer-richa/golang_basics/functions"
	"github.com/programmer-richa/golang_basics/helloworld"
	"github.com/programmer-richa/golang_basics/pointers"
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
	printSeperator("If ElseIf Else Example")

	controlflowstatements.IfElseIfElseExample()
	printSeperator("Switch With OptionalStatement Example")
	controlflowstatements.SwitchWithOptionalStatementExample()
	printSeperator("Switch Without Optional Statement Example")
	controlflowstatements.SwitchWithoutOptionalStatementExample()
	printSeperator("Switch Expression Example")
	controlflowstatements.SwitchExpressionExample()
	printSeperator("Switch Type Variable Example")
	controlflowstatements.SwitchTypeVariableExample()
	printSeperator("For Loop Example")
	controlflowstatements.ForLoopExample()
	printSeperator("For Loop With Break Example")
	controlflowstatements.ForLoopWithBreakExample()
	printSeperator("For Loop With Continue Example")
	controlflowstatements.ForLoopWithContinueExample()
	printSeperator("For Loop With Labled Continue Example")
	controlflowstatements.ForLoopWithLabledContinueExample()
	printSeperator("For Loop With GoTo Example")
	controlflowstatements.ForLoopWithGoToExample()
	printSeperator("For Infinite Loop Example")
	controlflowstatements.ForInfiniteLoopExample()
	printSeperator("For Loop Like While Example")
	controlflowstatements.ForLoopLikeWhileExample()
	printSeperator("For Loop On Range Example")
	controlflowstatements.ForLoopOnRangeExample()
	printSeperator("For Loop On String Example")
	controlflowstatements.ForLoopOnStringExample()
	printSeperator("For Loop On Map Example")
	controlflowstatements.ForLoopOnMapExample()
	printSeperator("For Loop on Channel Example")
	controlflowstatements.ForLoopOnChannelExample()

	printSeperator("One Dimentional Array Example With Zero Values")
	arrays.OneDimentionalArrayExampleWithZeroValues()
	printSeperator("One Dimentional Array Example")
	arrays.OneDimentionalArrayExample()
	printSeperator("Declaring And Initializing one Dimentional Array In One Line Example")
	arrays.DeclaringAndInitializingoneDimentionalArrayInOneLineExample()
	printSeperator("Declaring And Initializing MultiDimentional Array In One Line Example")
	arrays.DeclaringAndInitializingMultiDimentionalArrayInOneLineExample()
	printSeperator("One Dimentional Array With Size Determined By Values")
	arrays.OneDimentionalArrayWithSizeDeterminedByValues()
	printSeperator("Comparing Arrays")
	arrays.ComparingArrays()
	printSeperator("Copy Array By Value")
	arrays.CopyArrayByValue()
	printSeperator("Copy Array By Reference")
	arrays.CopyArrayByReference()

	printSeperator("Default Value Of Pointer")
	pointers.DefaultValueOfPointer()
	printSeperator("Pointer Declaration And Initialization")
	pointers.PointerDeclarationAndInitialization()
	printSeperator("Pointer Declaration And Initialization With Variable Reference")
	pointers.PointerDeclarationAndInitializationWithVariableReference()

	functionsModuleWorking()
}

func printSeperator(heading string) {
	println("\n\n", heading, "\n__________________________________________________________________________\n")
}

func functionsModuleWorking() {
	a := 10
	b := 20
	printSeperator("Max of 2 numbers")
	fmt.Printf("Max of %d and %d is %d\n", a, b, functions.Max(a, b))

	printSeperator("Swap 2 numbers by value")
	fmt.Println("In main before swapping x = ", a)
	fmt.Println("In main before swapping y = ", b)
	functions.SwapByValue(a, b)
	fmt.Println("In main after swapping x = ", a)
	fmt.Println("In main after swapping y = ", b)

	printSeperator("Swap 2 numbers by reference")
	fmt.Println("In main before swapping x = ", a)
	fmt.Println("In main before swapping y = ", b)
	functions.SwapByReference(&a, &b)
	fmt.Println("In main after swapping x = ", a)
	fmt.Println("In main after swapping y = ", b)

	printSeperator("Implode Function")
	line := functions.Implode(",", "I", "love", "go", "language", ".")
	fmt.Println("Line generated : ", line)

	printSeperator("Explode Function")
	tokens := functions.Explode(",", line)
	fmt.Println("Tokens generated : ", tokens)
	fmt.Println("Tokens length : ", len(tokens))

	printSeperator("Sum of 1D array Elements")
	arr := []int{10, 20, 30, 40, 50}
	total := functions.SumOf1DArray(arr)
	fmt.Println("Total of arr : ", arr, " is ", total)
}
