package main

import "fmt"

func main() {
	var Principal_amount int
	var Interest_rate float64
	var Number_of_years float64
	var Simple_interest float64
	fmt.Printf("Enter your Principal amount : ")
	fmt.Scanf("%d", &Principal_amount)
	fmt.Printf("Enter your Interest rate : ")
	fmt.Scanf("%f", &Interest_rate)
	fmt.Printf("Enter your Number of years : ")
	fmt.Scanf("%f", &Number_of_years)
	Simple_interest = (float64(Principal_amount) * Interest_rate * Number_of_years) / 100
	fmt.Printf("%f", Simple_interest)

}
