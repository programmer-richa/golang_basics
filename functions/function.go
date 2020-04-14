//Package functions implements examples of defining and calling functions.
package functions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag is used for indentation Purpose.
const level = 1
const sublevel = level + 1

// Array Modules Listing Constants
const (
	_ = iota
	maxConst
	swapConst
	swapreferenceConst
	implodeConst
	explodeConst
	arraySumConst
	variadicSumConst
	deferConst
	anonymousConst
	expressionConst
	returnConst
	callbackConst
	closureConst
	recursionConst
	exitConst
	maxConstLbl           = "Max of 2 numbers"
	swapConstLbl          = "Swap 2 numbers by value"
	swapreferenceConstLbl = "Swap 2 numbers by reference"
	implodeConstLbl       = "Join String"
	explodeConstLbl       = "Split String"
	arraySumConstLbl      = "Sum of 1D array Elements"
	variadicSumConstLbl   = "Sum of variable Elements"
	deferConstLbl         = "Deffering functions"
	anonymousConstLbl     = "Anonymous functions"
	expressionConstLbl    = "Function Expression i.e. referring anonymous function by variable name."
	returnConstLbl        = "Returning function from function"
	callbackConstLbl      = "Passing function as an argument. However, this approach is not recommended."
	closureConstLbl       = "Defining anonymous function inside a function to enclose value of a variable declared in enclosing function "
	recursionConstLbl     = "Factorial of a number using recursion"
	exitConstLbl          = "Exit"
)

//Block enables user to choose from the list of options available to test variety of functions implemented in this sub module.
func Block() {
	var a, b, choice = 10, 20, 0 // local variable
	tabs := utility.Tabs(sublevel)
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, maxConst, maxConstLbl)
		utility.Println(level, swapConst, swapConstLbl)
		utility.Println(level, swapreferenceConst, swapreferenceConstLbl)
		utility.Println(level, implodeConst, implodeConstLbl)
		utility.Println(level, explodeConst, explodeConstLbl)
		utility.Println(level, arraySumConst, arraySumConstLbl)
		utility.Println(level, variadicSumConst, variadicSumConstLbl)
		utility.Println(level, deferConst, deferConstLbl)
		utility.Println(level, anonymousConst, anonymousConstLbl)
		utility.Println(level, expressionConst, expressionConstLbl)
		utility.Println(level, returnConst, returnConstLbl)
		utility.Println(level, callbackConst, callbackConstLbl)
		utility.Println(level, closureConst, closureConstLbl)
		utility.Println(level, recursionConst, recursionConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case maxConst:
			utility.Printh(sublevel, maxConstLbl)
			fmt.Printf(tabs+"Max of %d and %d is %d\n", a, b, Max(a, b))
		case swapConst:
			utility.Printh(sublevel, swapConstLbl)
			utility.Println(sublevel, "In main before swapping x = ", a)
			utility.Println(sublevel, "In main before swapping y = ", b)
			SwapByValue(a, b)
			utility.Println(sublevel, "In main after swapping x = ", a)
			utility.Println(sublevel, "In main after swapping y = ", b)
		case swapreferenceConst:
			utility.Printh(sublevel, swapreferenceConstLbl)
			utility.Println(sublevel, "In main before swapping x = ", a)
			utility.Println(sublevel, "In main before swapping y = ", b)
			SwapByReference(&a, &b)
			utility.Println(sublevel, "In main after swapping x = ", a)
			utility.Println(sublevel, "In main after swapping y = ", b)
		case implodeConst:
			utility.Printh(sublevel, implodeConstLbl)
			line := Implode(",", "I", "love", "go", "language", ".")
			utility.Println(sublevel, "Line generated : ", line)
		case explodeConst:
			line := "I,love,go,language,."
			utility.Printh(sublevel, explodeConstLbl)
			tokens := Explode(",", line)
			utility.Println(sublevel, "Tokens generated : ", tokens)
			utility.Println(sublevel, "Tokens length : ", len(tokens))
		case arraySumConst:
			utility.Printh(sublevel, arraySumConstLbl)
			arr := []int{10, 20, 30, 40, 50}
			total := SumOf1DArray(arr)
			utility.Println(sublevel, "Total of arr : ", arr, " is ", total)
		case variadicSumConst:
			utility.Printh(sublevel, variadicSumConstLbl)
			total := Sum(10, 20, 30)
			utility.Println(sublevel, "Total of 10,20,30 : ", total)
			x := []int{1, 2, 3, 4, 5}
			// passing all values from slice x to Sum
			//unfurling the slice
			total = Sum(x...)
			utility.Println(sublevel, "Total of slice : ", x, " is ", total)
		case deferConst:
			utility.Printh(sublevel, deferConstLbl)
			Defer()
		case anonymousConst:
			Anonymous()
		case expressionConst:
			Expression()
		case returnConst:
			ReturnFunction()
		case callbackConst:
			CallBack()
		case closureConst:
			Closure()
		case recursionConst:
			Recursion()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Explode is a variadic Function that splits a string with given separator and returns string array.
func Explode(separator string, str string) []string {
	return strings.Split(str, separator)
}

//Implode is a variadic Functions that joins a string with given separator and returns string array.
func Implode(separator string, str ...string) string {
	return strings.Join(str, separator)
}

//SwapByValue demonstrates working of Pass By Value i.e separate copy of arguments is created in the memory,
// hence the changes made in the formal arguments are not reflected in the actual arguments.
func SwapByValue(x int, y int) {
	temp := x
	x = y
	y = temp
	utility.Println(sublevel, "In swapByValue x = ", x)
	utility.Println(sublevel, "In swapByValue y = ", y)
}

//SwapByReference demonstrates working of Pass By Reference i.e address of arguments passed to the function,
//  hence the changes made in the formal arguments are reflected in the actual arguments.
func SwapByReference(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
	utility.Println(sublevel, "In swapByValue x = ", *x)
	utility.Println(sublevel, "In swapByValue y = ", *y)
}

// Max finds greater of 2 number and returns it.
func Max(x int, y int) (z int) {

	if x > y {
		z = x
	}
	return z
}

// SumOf1DArray calculates sum of values in 1D array and returns total.
func SumOf1DArray(a []int) (total int) {
	sizeOnArray := len(a)
	for i := 0; i < sizeOnArray; i++ {
		total += a[i]
	}

	return total
}

// Sort1DArrayByValue sorts 1D array of data type int.
func Sort1DArrayByValue(a []int) {
	sort.Ints(a)
}

//Sum calculates total of variable int arguments
func Sum(a ...int) (total int) {
	for _, value := range a {
		total += value
	}
	return total
}

//Defer demnstrates working of defer keyword.
//Here closefile function is deferred by openFile function
func Defer() {
	//closeFile func executes once readFile func completes its execution
	defer CloseFile()
	//readFile func executes once openFile func completes its execution
	defer ReadFile()
	OpenFile()

}

// CloseFile closes a file
func CloseFile() {
	utility.Println(sublevel, "Closing a file")
}

// ReadFile reads data from a file
func ReadFile() {
	utility.Println(sublevel, "Reading a file")
}

// OpenFile opens a file
func OpenFile() {
	utility.Println(sublevel, "Opening a file")
}

//Anonymous demonstrates working of anonymous functions.
//Here anonymous function takes 2 arguments and returns sum of arguments.
func Anonymous() {
	utility.Println(sublevel, anonymousConstLbl)
	a, b := 10, 20
	// anonymous function
	x := func(a int, b int) (sum int) {
		sum = a + b
		return sum
	}(a, b) // here arguments are passed in function call
	utility.Println(sublevel, "Sum of", a, "and", b, "is", x)
}

//Expression demonstrates working of expression functions.
//Here expression function takes 2 arguments and returns max of arguments.
func Expression() {
	utility.Println(sublevel, expressionConstLbl)
	a, b := 10, 20
	// expression function
	max := func(a int, b int) (max int) {
		if a > b {
			max = a
		} else {
			max = b
		}
		return max
	}
	greater := max(a, b) // calling function using function expression
	utility.Println(sublevel, "Greater of", a, "and", b, "is", greater)
}

//ReturnFunction demonstrates working of functions that returns a function.
func ReturnFunction() {
	utility.Println(sublevel, expressionConstLbl)
	a, b := 10, 20
	multiply := 5
	// expression function
	table := TableSum(a, b) // returns function that could print table of sum of a and b
	data := table(multiply) // calling function using function expression
	utility.Println(sublevel, "Table of (10 + 20) * 5 \n", data)

	data = table(multiply + 5) // calling function using function expression
	utility.Println(sublevel, "Table of (10 + 20) * 10 \n", data)
}

// TableSum accepts 2 arguments of integer type and returns a function that returns a string that
// prints table of sum of arguments passed to TableSum function multiplied by the number specified
//  as an argument passed to the returned function
func TableSum(a int, b int) func(multiply int) string {
	sum := a + b
	tabs := utility.Tabs(sublevel)
	table := func(multiply int) string {
		//strconv.Itoa(int_value) -- converts int to string
		num := sum * multiply
		data := tabs + "Table of " + strconv.Itoa(num) + "\n"
		data += tabs + strings.Repeat("-", 20) + "\n"
		for i := 1; i < 11; i++ {
			data += fmt.Sprintf(tabs+"%d * %d = %d\n", num, i, num*i)
		}
		return data
	}
	return table
}

//CallBack demonstrates working of function as an argument.
func CallBack() {
	utility.Println(sublevel, callbackConstLbl)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := Sum(a...)
	utility.Println(sublevel, "Sum of numbers from 1 to 10 is", sum)
	// calculating sum of even values in the dataset by passing Sum function as first argument, and values as second argument
	sum = Even(Sum, a...)
	utility.Println(sublevel, "Sum of even numbers from 1 to 10 is", sum)
	// calculating sum of odd values in the dataset by passing Sum function as first argument, and values as second argument
	sum = Odd(Sum, a...)
	utility.Println(sublevel, "Sum of odd numbers from 1 to 10 is", sum)
}

// Even is used to filter even values from the passed values and calls
// func passed as an argument to return the result of the filtered data.
func Even(f func(a ...int) (total int), values ...int) (total int) {
	var data []int
	for _, v := range values {
		if v%2 == 0 {
			data = append(data, v)
		}
	}
	return f(data...)
}

// Odd is used to filter odd values from the passed values and calls
// func passed as an argument to return the result of the filtered data.
func Odd(f func(a ...int) (total int), values ...int) (total int) {
	var data []int
	for _, v := range values {
		if v%2 != 0 {
			data = append(data, v)
		}
	}
	return f(data...)
}

// Closure demonstrates working of closure functions
func Closure() {
	utility.Println(sublevel, closureConstLbl)
	seq1 := IntSeq()
	utility.Println(sublevel, "Sequence1 Initialized")
	utility.Println(sublevel, seq1())
	utility.Println(sublevel, seq1())
	utility.Println(sublevel, seq1())

	seq2 := IntSeq()
	utility.Println(sublevel, "Sequence2 Initialized")
	utility.Println(sublevel, seq2())
	utility.Println(sublevel, seq2())
	utility.Println(sublevel, seq2())
}

// IntSeq returns another function,
// which we define anonymously in the body of IntSeq.
// The returned function closes over the variable i to form a closure.
func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Recursion demonstrates working of calling a function inside itself
func Recursion() {
	utility.Println(sublevel, recursionConstLbl)
	fact := Factorial(5)
	utility.Println(sublevel, "Factorial of 5 =", fact)
}

// Factorial calculates factorial of a number
func Factorial(num int) int {
	if num == 1 {
		return num
	} else {
		return num * Factorial(num-1)
	}
}
