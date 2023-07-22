package main

import "fmt"

func main() {
	var number int
	flag := 0
	fmt.Printf("Enter the number :")
	fmt.Scanf("%d", &number)
	for i := 2; i < number; i++ {
		if number%i == 0 {
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Printf("Number is Prime")
	} else {
		fmt.Printf("Number is not Prime")
	}

}
