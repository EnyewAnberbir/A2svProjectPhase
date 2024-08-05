package main

import(
	"fmt"
	"errors"
)

func averageGrades(subjects map[string]float64)(float64, error){
	if len(subjects) == 0{
		return 0 , errors.New("no subjects to calculate average grade")
	}
	var total float64
	for _,grade := range subjects{
		total += grade
	}
	average := total / float64(len(subjects))
	return average, nil
}
func main(){
	var name string
	var subjectTotal int

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Enter number of Subjects:")
	fmt.Scanln(&subjectTotal)

	subjects := make(map[string]float64)

	for i := 0 ; i < subjectTotal; i++{
		var subjectName string
		var grade float64

		fmt.Printf("Enter name of Subject" , i+1)
		fmt.Scanln(&subjectName)

		fmt.Printf("Enter the grade", subjectName)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 100{
			fmt.Println("Enter correct Grade")
			i--
			continue
		}

		subjects[subjectName] = grade
	}
    averageGrades, err := averageGrades(subjects)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Student Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for subject, grade := range subjects {
		fmt.Printf("%s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrades)


}