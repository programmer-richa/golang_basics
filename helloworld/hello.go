package helloworld

/*
Here, import keyword is used to import packages in your program and fmt package is used to implement
formatted Input/Output with functions.
*/
// adding dependencies
import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
	"rsc.io/quote"
)

// Module Level Flag
const level = 1

//Greeting exported Functions must have comments as specified above
//Greeting prints Hello String
func Greeting() {
	utility.Println(level, "Hello world example")
	fmt.Println("Hello World")

}

/*
In the above lines of code we have:

func: It is used to create a function in Go language.
main: It is the main function in Go language, which doesn’t contain the parameter, doesn’t return anything, and call when you execute your program.
Println(): This method is present in fmt package and it is used to display “!… Hello World …!” string. Here, Println means Print line.
*/

//Hello Return hello world String
func Hello() string {
	return quote.Hello()

}
