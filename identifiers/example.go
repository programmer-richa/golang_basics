/*
In programming languages, identifiers are used for identification purposes.
 Or in other words, identifiers are the user-defined name of the program components.
 In Go language, an identifier can be a variable name, function name, constant, statement labels, package name, or types.

 Rules for Defining Identifiers: There are certain valid rules for defining a valid Go identifier.
These rules should be followed, otherwise, we will get a compile-time error.

The name of the identifier must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.
The name of the identifier should not start with a digit.
The name of the identifier is case sensitive.
Keywords is not allowed to use as an identifier name.
There is no limit on the length of the name of the identifier, but it is advisable to use an optimum length of 4 – 15 letters only.

Note:
The identifier represented by the underscore character(_) is known as a blank identifier. It is used as an anonymous placeholder instead of a regular identifier, and it has a special meaning in declarations, as an operand, and in assignments.
The identifier which is allowed to access it from another package is known as the exported identifier. The exported identifiers are those identifiers which obey the following conditions:
The first character of the exported identifier’s name should be in the Unicode upper case letter.
The identifier should be declared in the package block, or it is a variable name, or it is a method name.
The uniqueness of the identifiers means the identifier is unique from the other set of the identifiers available in your program, or in the package and they are not exported.
*/
package identifiers

import "github.com/programmer-richa/golang_basics/utility"

// Module Level Flag -- used for indentation Purpose
const level = 1

// Usage dummy identifier examples
func Usage() {
	utility.Printh(level, "Using Identifiers")
	var name = "Richa"
	utility.Println(level, "name =", name)
}
