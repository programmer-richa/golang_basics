package arrays

import "fmt"

//OneDimentionalArrayExampleWithZeroValues : As there is no concept of uninitiatized variables in go, all the array values are assigned zero(default) values according to the data type
func OneDimentionalArrayExampleWithZeroValues() {
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	// As the array elements are not assigned any values, the contain Zero(Default) Values as per data type
	for i := 0; i < subjectCount; i++ {
		fmt.Println("Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//OneDimentionalArrayWithSizeDeterminedByValues: Array length calculation according to the number of values passed
func OneDimentionalArrayWithSizeDeterminedByValues() {

	studentMarks := [...]int{10, 20, 30}
	//As 3 values are specified while initialization, the length of array is 3
	lengthOfArray := len(studentMarks)
	for i := 0; i < lengthOfArray; i++ {
		fmt.Println("Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//ComparingArrays: Comparing if two arrays are equal or not
func ComparingArrays() {

	studentMarks := [...]int{10, 20, 30}
	//Arrays are value type, so copy of array is created and assigned to new array
	studentMarks1 := studentMarks
	fmt.Println("Before modifing second array, comparison result: ", studentMarks == studentMarks1)
	studentMarks1[1] = 40
	fmt.Println("After modifing second array, comparison result: ", studentMarks == studentMarks1)

}

//CopyArrayByValue: Arrays are value type, so copy of array is created and assigned to new array
func CopyArrayByValue() {
	studentMarks := [...]int{10, 20, 30}
	studentMarks1 := studentMarks
	fmt.Println("Before modifing second array, studentMarks: ", studentMarks)
	fmt.Println("Before modifing second array, studentMarks1: ", studentMarks1)
	studentMarks1[1] = 40
	fmt.Println("After modifing second array, studentMarks: ", studentMarks)
	fmt.Println("After modifing second array, studentMarks1: ", studentMarks1)

}

//CopyArrayByReference: Arrays are value type, so address of array assigned to new array so that the changes made in either array are reflected in other
func CopyArrayByReference() {

	studentMarks := [...]int{10, 20, 30}
	studentMarks1 := &studentMarks
	fmt.Println("Before modifing second array, studentMarks: ", studentMarks)
	fmt.Println("Before modifing second array, studentMarks1: ", studentMarks1)
	studentMarks1[1] = 40
	fmt.Println("After modifing second array, studentMarks: ", studentMarks)
	fmt.Println("After modifing second array, studentMarks1: ", studentMarks1)

}

//OneDimentionalArrayExample : working of 1D Array
func OneDimentionalArrayExample() {
	const subjectCount int = 5
	var studentMarks [subjectCount]int
	for i := 0; i < subjectCount; i++ {
		studentMarks[i] = i * 10
	}

	for i := 0; i < subjectCount; i++ {
		fmt.Println("Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//DeclaringAndInitializingoneDimentionalArrayInOneLineExample : Declaring Multidimentional Array
func DeclaringAndInitializingoneDimentionalArrayInOneLineExample() {
	const subjectCount int = 5
	studentMarks := [subjectCount]int{10, 20, 30, 40, 50}

	for i := 0; i < subjectCount; i++ {
		fmt.Println("Marks in subject", i+1, ": ", studentMarks[i])
	}

}

//DeclaringAndInitializingMultiDimentionalArrayInOneLineExample : Declaring MultiDimetional Array
func DeclaringAndInitializingMultiDimentionalArrayInOneLineExample() {
	const studentCount int = 2
	const subjectCount int = 3
	studentMarks := [studentCount][subjectCount]int{{10, 20, 30},
		{40, 50, 60},
	}

	for student := 0; student < studentCount; student++ {
		fmt.Println("Marks of student", student+1, ": ")
		for subject := 0; subject < subjectCount; subject++ {
			fmt.Println("\tMarks in subject", subject+1, ": ", studentMarks[student][subject])
		}
	}

}
