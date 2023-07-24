package main

import "fmt"

func main() {
	var limit int
	fmt.Printf("Enter the limit :")
	fmt.Scanf("%d", &limit)
	a := make([][]int, limit)
	b := make([][]int, limit)
	sum := make([][]int, limit)
	for i := 0; i < limit; i++ {
		a[i] = make([]int, limit)
		b[i] = make([]int, limit)
		sum[i] = make([]int, limit)
	}
	getArray(limit, a, b)
	addArray(limit, a, b, sum)
	displayArray(limit, sum)
}
func getArray(n int, a [][]int, b [][]int) {
	fmt.Printf("Enetr the values of array 1:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	fmt.Printf("Enetr the values of array 2:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}
}
func addArray(n int, a [][]int, b [][]int, sum [][]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[i][j] = a[i][j] + b[i][j]
		}
	}
}
func displayArray(n int, sum [][]int) {
	fmt.Printf("Sum of Array :\n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", sum[i][j])
		}
		fmt.Println()
	}
}
