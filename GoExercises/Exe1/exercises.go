package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name        string
	StudentId   int
	DateOfBirth time.Time
}

func (s *Student) Age() int {
	now := time.Now()
	age := now.Year() - s.DateOfBirth.Year()
	if now.Month() < s.DateOfBirth.Month() ||
		(now.Month() == s.DateOfBirth.Month() && now.Day() < s.DateOfBirth.Day()) {
		age--
	}
	return age
}
func (s *Student) String() string {
	return (fmt.Sprintf(
		"Student Name: %s\nStudent ID: %d\nAge: %d\nDate of Birth: %s\n\n",
		s.Name,
		s.StudentId,
		s.Age(),
		s.DateOfBirth.Format("01-02-2006")))
}

func parseDate(date string) time.Time {
	dob, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return dob
}

func createStudent(name string, stdId int, date string) Student {
	return Student{Name: name, StudentId: stdId, DateOfBirth: parseDate(date)}
}
func main() {
	// dob, err := time.Parse("02-01-2006", "15-11-2004")
	// if err != nil {
	// 	fmt.Println("Error parsing date: ", err)
	// 	return
	// }

	// Create individual students
	Nahom := createStudent("Nahom", 101, "2003-04-15")
	Ola := createStudent("Ola", 102, "2003-04-15")
	Torbjørn := createStudent("Torbjørn", 103, "2003-04-15")
	Unni := createStudent("Unni", 104, "2003-04-15")
	Kari := createStudent("Kari", 105, "2003-04-15")
	// Create an array of students
	sArray := make([]Student, 5, 20)

	// Assign students to the array
	sArray[0] = Nahom
	sArray[1] = Ola
	sArray[2] = Torbjørn
	sArray[3] = Unni
	sArray[4] = Kari

	// var sArray2 []Student = make([]Student, 5)
	// _ = sArray2
	for _, student := range sArray {
		fmt.Println(student.String())
	}
	Erik := createStudent("Erik", 111, "2003-04-15")
	Maria := createStudent("Maria", 112, "2003-04-15")
	Anna := createStudent("Anna", 113, "2003-04-15")
	Jonas := createStudent("Jonas", 114, "2003-04-15")
	Sara := createStudent("Sara", 115, "2003-04-15")

	fmt.Println("after added students\n\n\n ")
	sArray = append(sArray, Erik, Maria, Anna, Jonas, Sara)

	for _, student := range sArray {
		fmt.Println(student.String())
	}
}
