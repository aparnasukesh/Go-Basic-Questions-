package main

import "fmt"

func main() {
	sum := 0
	var limit, i int
	fmt.Printf("Enter the limit")
	fmt.Scanf("%d", &limit)
	for i = 1; i <= limit; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Printf("Sum of Odd Numbers = %d", sum)
}
