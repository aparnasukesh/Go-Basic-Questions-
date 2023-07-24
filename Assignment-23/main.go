package main

import "fmt"

func main() {
	var limit int
	fmt.Printf("enter the array limit:")
	fmt.Scanf("%d", &limit)
	a := make([][]int, limit)
	for i := 0; i < limit; i++ {
		a[i] = make([]int, limit)
	}
	getarray(limit, a)
	displayarray(limit, a)
}
func getarray(limit int, a [][]int) {
	fmt.Printf("enter the array values:")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
}
func displayarray(limit int, a [][]int) {
	fmt.Printf(" array values are:\n")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}

}
