package main

import "fmt"

func main() {
	var arr1 [10][10]int
	var arr2 [10][10]int
	var Sum [10][10]int
	var limit, i int
	fmt.Printf("Enter the limit")
	fmt.Scanf("%d", &limit)
	fmt.Printf("Enter the values of array1:")
	for i = 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}
	fmt.Printf("\nEnter the values of array2:")
	for i = 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
	}
	fmt.Printf("\nArray 1 :\n")
	for i = 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%d ", arr1[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("\nArray 2 :\n")
	for i = 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%d ", arr2[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Sum Array :\n")
	for i = 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			Sum[i][j] = arr1[i][j] + arr2[i][j]
			fmt.Printf("%d ", Sum[i][j])
		}
		fmt.Println()

	}

}
