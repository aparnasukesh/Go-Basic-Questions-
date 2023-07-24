package main

import "fmt"

func main() {
	var written_test, lab_exams, assignments float64
	fmt.Printf("Enter your marks for written test :")
	fmt.Scanf("%f", &written_test)
	fmt.Printf("Enter your marks for lab exam :")
	fmt.Scanf("%f", &lab_exams)
	fmt.Printf("Enter your marks for assignments :")
	fmt.Scanf("%f", &assignments)
	var over_all_grade = (written_test*70)/100 + (lab_exams*20)/100 + (assignments*10)/100
	fmt.Printf("Over All Grade = %f", over_all_grade)
}
