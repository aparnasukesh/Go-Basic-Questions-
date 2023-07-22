package main

import "fmt"

func main() {
	var count, i, limit int
	var arr [10]int
	fmt.Printf("Enter the limit")
	fmt.Scanf("%d", &limit)
	fmt.Printf("Enter the values of array : ")
	for i = 0; i < limit; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i = 0; i < limit; i++ {
		if arr[i]%2 == 0 {
			count = count + 1
		}
	}
	fmt.Printf("Number of even numbers in the array : %d", count)
}
