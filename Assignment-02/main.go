package main

import "fmt"

func main() {
	var number1 int
	var number2 float64
	var sum float64

	fmt.Printf("Enter two numbers:")
	fmt.Scanf("%d%f", &number1, &number2)
	sum = float64(number1) + number2
	fmt.Printf("Sum = %f", sum)

}
