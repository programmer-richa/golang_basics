// Package strings demontrats working of strings
package strings

import "github.com/programmer-richa/golang_basics/utility"

// Module Level Flag is used for indentation purpose.
const level = 1
const sublevel = level + 1

// ToByteArray   coverts string to byte array.
func ToByteArray() {
	utility.Printh(sublevel, "String to []byte conversion")
	str := "Hello, world"
	byteArr := []byte(str)
	utility.Println(sublevel, "Value of str is %v having length %d and Type of str : %T", str, len(str), str)
	utility.Println(sublevel, "Value of byteArr is %v having length %d and Type of byteArr : %T", byteArr, len(byteArr), byteArr)

	utility.Println(sublevel, "Printing all unicode characters in string")
	for key, value := range str {
		utility.Println(sublevel, "%s is stored at Index :%d having binary code %b, unicode value :%d and hexcode value %#x", string(str[key]), key, value, value, value)
	}
}
