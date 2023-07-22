package main

import "fmt"

func main() {
	var mark float64
	fmt.Printf("Enter your mark : ")
	fmt.Scanf("%f", &mark)
	if mark > 90 {
		fmt.Println("Grade : A")
	} else if mark >= 80 {
		fmt.Println("Grade : B")
	} else if mark >= 70 {
		fmt.Println("Grade : C")
	} else if mark >= 60 {
		fmt.Println("Garde : D")
	} else if mark >= 50 {
		fmt.Println("Grade : E")
	} else {
		fmt.Println("Failed !")
	}

}
