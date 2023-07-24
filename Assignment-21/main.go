package main

import "fmt"

func main() {
	var limit int
	fmt.Printf("Enter the limit")
	fmt.Scan(&limit)
	a := make([]int, limit)
	fmt.Printf("Enter the values of array:")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println("Array :")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Print("\nAfter multiplying adjacent elements in the, array :")
	for i := 0; i < limit-1; i++ {
		a[i] = a[i] * a[i+1]
		fmt.Printf("%d ", a[i])
	}
}
