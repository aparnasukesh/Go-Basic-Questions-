package main

import "fmt"

func main() {
	var limit int
	var temp = 1
	fmt.Printf("Enter the limit")
	fmt.Scan(&limit)
	for i := 1; i <= limit; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", temp)
			temp++
		}
		fmt.Println()
	}
}
