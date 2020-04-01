package controlflowstatements

import "fmt"

// IfElseIfElseExample : If elseif statement
func IfElseIfElseExample() {
	var a int = 10
	var b int = 20
	if a > b {
		fmt.Println(a, "is greater than", b)
	} else if a < b {
		fmt.Println(a, "is smaller than", b)
	} else {
		fmt.Println("Both are equal")
	}
}

// SwitchWithOptionalStatementExample : switch statement
func SwitchWithOptionalStatementExample() {
	// Switch statement with both
	// optional statement, i.e, day:=4
	// and expression, i.e, day
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}
}

// SwitchWithoutOptionalStatementExample : switch without choice variable
func SwitchWithoutOptionalStatementExample() {
	var value int = 2

	// Switch statement without an
	// optional statement and
	// expression
	switch {
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Bonjour")
	case value == 3:
		fmt.Println("Namstay")
	default:
		fmt.Println("Invalid")
	}
}

// SwitchExpressionExample : switch with multiple values in case statement
func SwitchExpressionExample() {
	var value string = "three"

	// Switch statement without default statement
	// Multiple values in case statement
	switch value {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}
}

// SwitchTypeVariableExample : switch based on data type of variable
func SwitchTypeVariableExample() {
	var value interface{}
	switch q := true; value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}
}

// ForLoopExample : Simple Loop
func ForLoopExample() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}

// ForLoopWithBreakExample : break statement example
func ForLoopWithBreakExample() {
	for i := 1; i < 11; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

// ForLoopWithContinueExample : continue statement example
func ForLoopWithContinueExample() {
	for i := 1; i < 11; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}

// ForLoopWithLabledContinueExample : continue statement example
func ForLoopWithLabledContinueExample() {

Lable1:
	for i := 1; i < 5; i++ {
		fmt.Println("Value of i: ", i)
		for j := 1; j <= 5; j++ {
			if j == 3 {
				continue Lable1
			}
			fmt.Println("Value of j: ", j)
		}

	}
}

// ForLoopWithGoToExample : goto statement example
func ForLoopWithGoToExample() {

	var x int = 0

	// for loop work as a while loop
Lable1:
	for x < 8 {
		if x == 5 {

			// using goto statement
			x = x + 1
			goto Lable1
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}
}

// ForInfiniteLoopExample : Infinite Loop
func ForInfiniteLoopExample() {
	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}

// ForLoopLikeWhileExample : For loop with only condition statement
func ForLoopLikeWhileExample() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

// ForLoopOnRangeExample : Loops over a array or range of values
func ForLoopOnRangeExample() {
	// Here rvariable is a array
	rvariable := []string{"Ram", "Sham", "Month"}

	// key and value stores the value of rvariable
	// key store index number of individual string and
	// value store individual string of the given array
	for key, value := range rvariable {
		fmt.Println(key, value)
	}
}

// ForLoopOnStringExample : Loops over a string
func ForLoopOnStringExample() {
	// String as a range in the for loop
	for i, j := range "Richa" {
		fmt.Printf("The index number of %c is %d\n", j, i)
	}
}

// ForLoopOnMapExample : Loops over a map
func ForLoopOnMapExample() {
	// using maps
	mmap := map[int]string{
		22: "Ram",
		33: "Mohan",
		44: "Sohan",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
}

// ForLoopOnChannelExample : Loops over channel values
func ForLoopOnChannelExample() {
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
		fmt.Println(i)
	}
}
