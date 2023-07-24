package main

import "fmt"

func main() {
	var choice int
	var num1, num2 int
	fmt.Printf("Enter two numbers :")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Enter your choice :")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		addition(num1, num2)
	case 2:
		subtraction(num1, num2)
	case 3:
		multiplication(num1, num2)
	case 4:
		division(num1, num2)
	default:
		fmt.Println("Invalid entry..")
	}
}
func addition(n1, n2 int) {
	var result int
	result = n1 + n2
	fmt.Printf("Result = %d", result)
}
func subtraction(n1, n2 int) {
	var result int
	result = n1 - n2
	fmt.Printf("Result = %d", result)
}
func multiplication(n1, n2 int) {
	var result int
	result = n1 * n2
	fmt.Printf("Result = %d", result)
}
func division(n1, n2 int) {
	var result int
	result = n1 / n2
	fmt.Printf("Result = %d", result)
}
