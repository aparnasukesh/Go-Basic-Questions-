package main

import "fmt"

func main() {
	var mark int
	fmt.Printf("Enter your mark")
	fmt.Scanf("%d", &mark)
	if mark < 50 {
		fmt.Printf("Student failed")
	} else {
		fmt.Printf("Student passed ")
	}

}
