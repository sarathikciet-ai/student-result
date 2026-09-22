package main

import (
	"fmt"

	"student-result/models"
	"student-result/result"
)

func main() {
	var student models.Student
	var numberOfSubjects int

	fmt.Println("===== Student Result Processing System =====")
	fmt.Println("Welcome to the Application")

	fmt.Print("Enter Student Name: ")
	fmt.Scanln(&student.Name)

	fmt.Print("Enter Roll Number: ")
	fmt.Scanln(&student.RollNo)

	fmt.Print("Enter Number of Subjects: ")
	fmt.Scan(&numberOfSubjects)

	student.Marks = make([]float64, numberOfSubjects)

	for i := 0; i < numberOfSubjects; i++ {
		fmt.Printf("Enter marks for Subject %d: ", i+1)
		fmt.Scan(&student.Marks[i])
	}

	total := result.CalculateTotal(student.Marks)
	average := result.CalculateAverage(student.Marks)
	grade := result.CalculateGrade(average)
	status := result.CalculateResult(student.Marks)

	fmt.Println()
	fmt.Println("========== STUDENT RESULT ==========")
	fmt.Println("Name       :", student.Name)
	fmt.Println("Roll Number:", student.RollNo)
	fmt.Println("Marks      :", student.Marks)
	fmt.Printf("Total      : %.2f\n", total)
	fmt.Printf("Average    : %.2f\n", average)
	fmt.Println("Grade      :", grade)
	fmt.Println("Result     :", status)
	fmt.Println("====================================")
}
