package main

import (
	"errors"
	"fmt"
)

func averageGrades(subjects map[string]float64) (float64, error) {
	if len(subjects) == 0 {
		return 0, errors.New("no subjects to calculate average grade")
	}
	var total float64
	for _, grade := range subjects {
		total += grade
	}
	average := total / float64(len(subjects))
	return average, nil
}

func main() {
	var name string
	var subjectTotal int
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Enter number of subjects:")
	fmt.Scanln(&subjectTotal)

	subjects := make(map[string]float64)
	for i := 0; i < subjectTotal; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("Enter name of subject %d: ", i+1)
		fmt.Scanln(&subjectName)
		fmt.Printf("Enter the grade for %s: ", subjectName)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Enter a valid grade between 0 and 100.")
			i--
			continue
		}

		subjects[subjectName] = grade
	}

	averageGrade, err := averageGrades(subjects)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Student Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for subject, grade := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("AveGrade: %.2f\n", averageGrade)
}
