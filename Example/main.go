package main

import "fmt"

func main() {
	var i, number, limit, result int
	fmt.Printf("Enter the number")
	fmt.Scanf("%d", &number)
	fmt.Printf("Enter the limit")
	fmt.Scanf("%d", &limit)
	for i = 1; i <= limit; i++ {
		result = i * number
		fmt.Printf("%d X %d = %d", i, number, result)
		fmt.Println()
	}
}
