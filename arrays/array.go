package arrays

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag -- used for indentation Purpose
const level = 1
const sublevel = level + 1

// Array Modules Listing Constants
const (
	_ = iota
	zeroConst
	oneDimentionalConst
	initializeConst
	multidimentionalConst
	autoSizeConst
	compareConst
	copyValueConst
	copyReferenceConst
	exitConst
	zeroConstLbl             = "One Dimentional Array Example With Zero Values"
	oneDimentionalConstLbl   = "One Dimentional Array Example"
	initializeConstLbl       = "Declaring And Initializing one Dimentional Array In One Line Example"
	multidimentionalConstLbl = "Declaring And Initializing MultiDimentional Array In One Line Example"
	autoSizeConstLbl         = "One Dimentional Array With Size Determined By Values"
	compareConstLbl          = "Comparing Arrays"
	copyValueConstLbl        = "Copy Array By Value"
	copyReferenceConstLbl    = "Copy Array By Reference"
	exitConstLbl             = "Exit"
)

//Block Works according to choice of user
func Block() {
	var choice int
	//Module Level Option Display
	for choice != exitConst {
		utility.Printh(level, "Choose From Options below :")
		utility.Println(level, zeroConst, zeroConstLbl)
		utility.Println(level, oneDimentionalConst, oneDimentionalConstLbl)
		utility.Println(level, initializeConst, initializeConstLbl)
		utility.Println(level, multidimentionalConst, multidimentionalConstLbl)
		utility.Println(level, autoSizeConst, autoSizeConstLbl)
		utility.Println(level, compareConst, compareConstLbl)
		utility.Println(level, copyValueConst, copyValueConstLbl)
		utility.Println(level, copyReferenceConst, copyReferenceConstLbl)
		utility.Println(level, exitConst, exitConstLbl)
		utility.Print(level, "Enter Your choice :")

		//Get user module choice
		fmt.Scan(&choice)
		// Execute The Program according to user choice

		switch choice {
		case zeroConst:
			Zero()
		case oneDimentionalConst:
			OneDimentional()
		case initializeConst:
			Initialize()
		case multidimentionalConst:
			Multidimentional()
		case autoSizeConst:
			AutoSize()
		case compareConst:
			Compare()
		case copyValueConst:
			CopyValue()
		case copyReferenceConst:
			CopyReference()
		case exitConst:
			utility.Println(level, exitConstLbl)
		default:
			utility.Println(level, "Invalid Choice")
		}
	}

}

//Zero : As there is no concept of uninitiatized variables in go, all the array values are assigned zero(default) values according to the data type
func Zero() {
	utility.Printh(sublevel, zeroConstLbl)
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	// As the array elements are not assigned any values, the contain Zero(Default) Values as per data type
	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//AutoSize: Array length calculation according to the number of values passed
func AutoSize() {
	utility.Printh(sublevel, autoSizeConstLbl)
	studentMarks := [...]int{10, 20, 30}
	//As 3 values are specified while initialization, the length of array is 3
	lengthOfArray := len(studentMarks)
	for i := 0; i < lengthOfArray; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//Compare: Comparing if two arrays are equal or not
func Compare() {
	utility.Printh(sublevel, compareConstLbl)
	studentMarks := [...]int{10, 20, 30}
	//Arrays are value type, so copy of array is created and assigned to new array
	studentMarks1 := studentMarks
	utility.Println(sublevel, "Before modifing second array, comparison result: ", studentMarks == studentMarks1)
	studentMarks1[1] = 40
	utility.Println(sublevel, "After modifing second array, comparison result: ", studentMarks == studentMarks1)

}

//CopyValue: Arrays are value type, so copy of array is created and assigned to new array
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

//CopyReference: Arrays are value type, so address of array assigned to new array so that the changes made in either array are reflected in other
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

//OneDimentional : working of 1D Array
func OneDimentional() {
	utility.Printh(sublevel, oneDimentionalConstLbl)
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	for i := 0; i < subjectCount; i++ {
		studentMarks[i] = i * 10
	}

	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//Initialize : Declaring Multidimentional Array
func Initialize() {
	utility.Printh(sublevel, initializeConstLbl)
	const subjectCount int = 5
	studentMarks := [subjectCount]int{10, 20, 30, 40, 50}

	for i := 0; i < subjectCount; i++ {
		utility.Println(sublevel, "Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//Multidimentional : Declaring MultiDimetional Array
func Multidimentional() {
	utility.Printh(sublevel, multidimentionalConstLbl)
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
