package main

import "fmt"

func main() {
	var Day int
	fmt.Printf("Enter your choice :")
	fmt.Scanf("%d", &Day)
	switch Day {
	case 1:
		fmt.Println("Sunady")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid entry...")
	}
}
