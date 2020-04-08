/*
Keywords or Reserved words are the words in a language that are used for some internal process or represent some predefined actions.
These words are therefore not allowed to use as an identifier. Doing this will result in a compile-time error.
*/
// Go program to illustrate the
// use of keywords
package keywords

import "github.com/programmer-richa/golang_basics/utility"

// Module Level Flag -- used for indentation Purpose
const level = 1

//Usage : package, import, func, var are keywords
func Usage() {
	utility.Printh(level, "Using Keywords")
	// Here, a is a valid identifier
	var a = "Richa"

	utility.Println(level, a)

	// Here, the default is an
	// illegal identifier and
	// compiler will throw an error
	// var default = "GFG"
}
