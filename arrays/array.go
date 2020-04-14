//Package arrays implements examples to demonstrate working of arrays.
// However it is recommended to use slice instead of array.
package arrays

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag is used for indentation Purpose
const level = 1
const sublevel = level + 1

// Array Modules Listing Constants
const (
	_ = iota
	zeroConst
	studentMarksConst
	storeMarksConst
	classMarksConst
	autoSizeConst
	compareConst
	copyValueConst
	copyReferenceConst
	loop1DConst
	loop2DConst
	exitConst
	zeroConstLbl          = "Zero Values of Array"
	studentMarksConstLbl  = "1D array to store marks obtained by a student in multiple subjects"
	storeMarksConstLbl    = "1D array to store marks obtained by a student in multiple subjects at declaration"
	classMarksConstLbl    = "2D array to store marks of multiple students in diffent subjects"
	autoSizeConstLbl      = "Array size calculation according to the number of values passed"
	compareConstLbl       = "Comparing Arrays"
	copyValueConstLbl     = "Copy Array By Value"
	copyReferenceConstLbl = "Copy Array By Reference"
	loop1DConstLbl        = "Prints index value pair of values in 1 D array using range"
	loop2DConstLbl        = "Prints index value pair of values in 2 D array using range"
	exitConstLbl          = "Exit"
)

//Block enables user to choose from the list of options available to test variety of array examples implemented in this sub module.
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, zeroConst, zeroConstLbl)
		utility.Println(level, studentMarksConst, studentMarksConstLbl)
		utility.Println(level, storeMarksConst, storeMarksConstLbl)
		utility.Println(level, classMarksConst, classMarksConstLbl)
		utility.Println(level, autoSizeConst, autoSizeConstLbl)
		utility.Println(level, compareConst, compareConstLbl)
		utility.Println(level, copyValueConst, copyValueConstLbl)
		utility.Println(level, copyReferenceConst, copyReferenceConstLbl)
		utility.Println(level, loop1DConst, loop1DConstLbl)
		utility.Println(level, loop2DConst, loop2DConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case zeroConst:
			Zero()
		case studentMarksConst:
			StudentMarks()
		case storeMarksConst:
			StoreMarks()
		case classMarksConst:
			ClassMarks()
		case autoSizeConst:
			AutoSize()
		case compareConst:
			Compare()
		case copyValueConst:
			CopyValue()
		case copyReferenceConst:
			CopyReference()
		case loop1DConst:
			Loop1D()
		case loop2DConst:
			Loop2D()
		case exitConst:
			utility.Println(level, exitConstLbl)

		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Zero values of variables are displayed in it. Since there is no concept of uninitiatized variables in go,
// all the array values are assigned zero(default) values according to the data type
func Zero() {
	utility.Printh(sublevel, zeroConstLbl)
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	// As the array elements are not assigned any values, the contain Zero(Default) Values as per data type
	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//AutoSize demonstrates the array length calculation according to the number of values passed.
func AutoSize() {
	utility.Printh(sublevel, autoSizeConstLbl)
	studentMarks := [...]int{10, 20, 30}
	//As 3 values are specified while initialization, the length of array is 3
	lengthOfArray := len(studentMarks)
	for i := 0; i < lengthOfArray; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//Compare demonstrates how arrays can be checked for equality
func Compare() {
	utility.Printh(sublevel, compareConstLbl)
	studentMarks := [...]int{10, 20, 30}
	//Arrays are value type, so copy of array is created and assigned to new array
	studentMarks1 := studentMarks
	utility.Println(sublevel, "Before modifing second array, comparison result: ", studentMarks == studentMarks1)
	studentMarks1[1] = 40
	utility.Println(sublevel, "After modifing second array, comparison result: ", studentMarks == studentMarks1)

}

//CopyValue demonstrates that arrays are of value type, therefore, a copy of array is created and assigned to new array.
func CopyValue() {
	utility.Printh(sublevel, copyValueConstLbl)
	studentMarks := [...]int{10, 20, 30}
	studentMarks1 := studentMarks
	utility.Println(sublevel, "Before modifing second array, studentMarks: ", studentMarks)
	utility.Println(sublevel, "Before modifing second array, studentMarks1: ", studentMarks1)
	studentMarks1[1] = 40
	utility.Println(sublevel, "After modifing second array, studentMarks: ", studentMarks)
	utility.Println(sublevel, "After modifing second array, studentMarks1: ", studentMarks1)

}

//CopyReference demonstrates that arrays are value type, therefore, address of array assigned to new array
// so that the changes made in either array are reflected in other
func CopyReference() {
	utility.Printh(sublevel, copyReferenceConstLbl)
	studentMarks := [...]int{10, 20, 30}
	studentMarks1 := &studentMarks
	utility.Println(sublevel, "Before modifing second array, studentMarks: ", studentMarks)
	utility.Println(sublevel, "Before modifing second array, studentMarks1: ", studentMarks1)
	studentMarks1[1] = 40
	utility.Println(sublevel, "After modifing second array, studentMarks: ", studentMarks)
	utility.Println(sublevel, "After modifing second array, studentMarks1: ", studentMarks1)

}

//StudentMarks implements 1D array to store marks obtained by a student in multiple subjects.
func StudentMarks() {
	utility.Printh(sublevel, storeMarksConstLbl)
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	for i := 0; i < subjectCount; i++ {
		studentMarks[i] = i * 10
	}

	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//StoreMarks  implements 1D array to store marks obtained by a student in multiple subjects at time of declaration.
func StoreMarks() {
	utility.Printh(sublevel, storeMarksConstLbl)
	const subjectCount int = 5
	studentMarks := [subjectCount]int{10, 20, 30, 40, 50}

	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//ClassMarks  implements 2D array to store marks of multiple students in diffent subjects.
func ClassMarks() {
	utility.Printh(sublevel, classMarksConstLbl)
	const studentCount int = 2
	const subjectCount int = 3
	studentMarks := [studentCount][subjectCount]int{{10, 20, 30},
		{40, 50, 60},
	}

	for student := 0; student < studentCount; student++ {
		utility.Println(sublevel, "Marks of student", student+1, ": ")
		for subject := 0; subject < subjectCount; subject++ {
			utility.Println(sublevel, "\tMarks in subject", subject+1, ": ", studentMarks[student][subject])
		}
	}

}

//Loop1D prints index value pair of values in array using range.
func Loop1D() {
	utility.Printh(sublevel, loop1DConstLbl)
	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		utility.Println(sublevel, "Index : ", index, "Value:", value)
	}
}

//Loop2D prints index value pair of values in 2 D array using range.
func Loop2D() {
	utility.Printh(sublevel, loop2DConstLbl)
	arr := [5][2]int{{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}
	for row, a := range arr {
		utility.Println(sublevel, "Row : ", row)
		for col, value := range a {
			utility.Println(sublevel, "Col : ", col, "value:", value)
		}
	}
}
