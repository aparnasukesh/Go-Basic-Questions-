package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]int
	var limit, i, temp int
	fmt.Printf("Enter the limit:")
	fmt.Scanf("%d", &limit)
	fmt.Printf("Enter the values of array1:")
	for i = 0; i < limit; i++ {
		fmt.Scanf("%d", &arr1[i])
	}
	fmt.Printf("Enter the values of array2:")
	for i = 0; i < limit; i++ {
		fmt.Scanf("%d", &arr2[i])
	}
	fmt.Printf("\nArray 1 : ")
	for i = 0; i < limit; i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Printf("\nArray 2:")
	for i = 0; i < limit; i++ {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Printf("\nAfter interchanging the values of array:")
	for i = 0; i < limit; i++ {
		for i = 0; i < limit; i++ {
			temp = arr1[i]
			arr1[i] = arr2[i]
			arr2[i] = temp
		}
	}
	fmt.Printf("\nArray 1 : ")
	for i = 0; i < limit; i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Printf("\nArray 2:")
	for i = 0; i < limit; i++ {
		fmt.Printf("%d ", arr2[i])
	}

}
