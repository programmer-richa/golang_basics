package exercise

import (
	"github.com/programmer-richa/golang_basics/utility"
)

// Student has attributes : rollno, name and marks
type Student struct {
	rollno int
	marks  float64
	name   string
}

// StudentFunc defines function set of Student struct
type StudentFunc interface {
	Print()
}

// Print displays student information
// Attaching Print method with Student struct
func (student Student) Print() {
	utility.Println(sublevel, "Roll No : ", student.Rollno())
	utility.Println(sublevel, "Name : ", student.Name())
	utility.Println(sublevel, "Marks : ", student.Marks())
}

// SetMarks assigns value to student marks.
// Attaching SetMarks method with Student.
// Here Student variable is passed ab reference so that the value is modified in the actual argument.
func (student *Student) SetMarks(marks float64) {
	student.marks = marks
}

// SetRollno assigns value to student rollno
// Attaching SetRollno method with Student struct
// Here Student variable is passed ab reference so that the value is modified in the actual argument.
func (student *Student) SetRollno(rollno int) {
	student.rollno = rollno
}

// SetName assigns value to student name
// Attaching SetName method with Student struct
// Here Student variable is passed ab reference so that the value is modified in the actual argument.
func (student *Student) SetName(name string) {
	student.name = name
}

// Marks returns student marks
// Attaching Marks method with Student struct
func (student Student) Marks() (marks float64) {
	return student.marks
}

// Rollno returns student rollno
// Attaching Rollno method with Student struct
func (student Student) Rollno() (rollno int) {
	return student.rollno
}

// Name returns student name
// Attaching Name method with Student struct
func (student Student) Name() (name string) {
	return student.name
}

// Show function accepts argument passed by value as well as reference but not a pointer variable because StudentFunc interface is specified as argument
func Show(stu StudentFunc) {
	stu.Print()
}

// Print function accepts a pointer to student variable
func Print(stu *Student) {
	(*stu).Print()
}

// Get function accepts a student variable
func Get(stu Student) {
	stu.Print()
}

// Struct demonstrates working of pointers and reference
func Struct() {
	var student = Student{}
	student.SetRollno(101)
	student.SetName("Richa")
	student.SetMarks(90)
	student.Print()
	utility.Println(sublevel, "Passing struct by value to func accepting interface as an argument")
	Show(student)
	utility.Println(sublevel, "Passing struct by reference to func accepting interface as an argument")
	Show(&student)
	// The following line generates error
	// Show(*student)
	utility.Println(sublevel, "Passing pointer variable")
	stuptr := &student
	Print(stuptr)

	utility.Println(sublevel, "Passing struct by value to func accepting student struct as an argument")
	Get(student)
	// The following line generates error
	// Get(&student)

}
