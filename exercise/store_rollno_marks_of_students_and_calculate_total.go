package exercise

import (
	"fmt"

	"github.com/programmer-richa/golang_basics/utility"
)

// dynamicMarks demonstrates working of slice according to number of students,subjects specified by user.
func dynamicMarks() {
	var studentsCount int
	utility.Println(sublevel, "Enter number Of Students")
	fmt.Scan(&studentsCount)
	// initialize slice according to number of students specified by user
	var students = make([][]int, studentsCount)
	for i := 0; i < studentsCount; i++ {
		var rollNumber, subjectsCount, total int = 0, 0, 0
		utility.Print(sublevel, "Enter roll number Of subjects of student ", i+1, ": ")
		fmt.Scan(&rollNumber)
		utility.Print(sublevel, "Enter number Of subjects of student ", i+1, ": ")
		fmt.Scan(&subjectsCount)
		// initialize slice according to number of subjects specified by user for the specific user, here 1st index is saved for roll number and last index for storing total marks of student
		students[i] = make([]int, subjectsCount+2)
		// stores roll number on first index
		students[i][0] = rollNumber
		for j := 1; j <= subjectsCount; j++ {
			utility.Print(sublevel, "Enter marks in subject ", j, " : ")
			// stores marks of student in given subject
			fmt.Scan(&students[i][j])
			total += students[i][j]
		}
		// stores total marks of student on last index
		students[i][subjectsCount+1] = total
	}

	for i := 0; i < studentsCount; i++ {
		subjectsCount := len(students[i])
		lastIndex := subjectsCount - 1
		utility.Println(sublevel, "Record of student", i+1, ":")
		utility.Println(sublevel, "Roll Number", ":", students[i][0])
		for j := 1; j < subjectsCount-1; j++ {
			utility.Println(sublevel, "Marks in subject", j, ":", students[i][j])
		}
		utility.Println(sublevel, "Total Marks", ":", students[i][lastIndex], "\n\n")
	}

}
