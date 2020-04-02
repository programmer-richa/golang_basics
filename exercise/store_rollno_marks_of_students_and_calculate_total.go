package exercise

import "fmt"

// GetStudentMarksDetails demonstrates working of slice according to number of students,subjects specified by user
func GetStudentMarksDetails() {
	var studentsCount int
	fmt.Print("Enter number Of Students")
	fmt.Scanln(&studentsCount)
	// initialize slice according to number of students specified by user
	var students = make([][]int, studentsCount)
	for i := 0; i < studentsCount; i++ {
		var rollNumber, subjectsCount, total int = 0, 0, 0
		fmt.Print("Enter roll number Of subjects of student ", i+1, ": ")
		fmt.Scanln(&rollNumber)
		fmt.Print("Enter number Of subjects of student ", i+1, ": ")
		fmt.Scanln(&subjectsCount)
		// initialize slice according to number of subjects specified by user for the specific user, here 1st index is saved for roll number and last index for storing total marks of student
		students[i] = make([]int, subjectsCount+2)
		// stores roll number on first index
		students[i][0] = rollNumber
		for j := 1; j <= subjectsCount; j++ {
			fmt.Print("Enter marks in subject ", j, " : ")
			// stores marks of student in given subject
			fmt.Scanln(&students[i][j])
			total += students[i][j]
		}
		// stores total marks of student on last index
		students[i][subjectsCount+1] = total
	}

	for i := 0; i < studentsCount; i++ {
		subjectsCount := len(students[i])
		lastIndex := subjectsCount - 1
		fmt.Println("Record of student", i+1, ":")
		fmt.Println("Roll Number", ":", students[i][0])
		for j := 1; j < subjectsCount-1; j++ {
			fmt.Println("Marks in subject", j, ":", students[i][j])
		}
		fmt.Println("Total Marks", ":", students[i][lastIndex], "\n\n")
	}

}
