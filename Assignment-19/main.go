package main

import "fmt"

func main() {
	var annual_income float64
	fmt.Printf("Enter your annual income :")
	fmt.Scan(&annual_income)
	var income_tax_amount float64
	if annual_income <= 250000 {
		fmt.Printf("No Income Tax")
	} else if annual_income <= 500000 {
		income_tax_amount = (annual_income * 5) / 100
		fmt.Printf("Income Tax Amount = %f", income_tax_amount)
	} else if annual_income <= 1000000 {
		income_tax_amount = (annual_income * 20) / 100
		fmt.Printf("Income Tax Amount = %f", income_tax_amount)
	} else if annual_income <= 5000000 {
		income_tax_amount = (annual_income * 30) / 100
		fmt.Printf("Income Tax Amount = %f", income_tax_amount)
	}
}
