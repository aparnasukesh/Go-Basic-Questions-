package main

import "fmt"

func main() {
	var limit int
	fmt.Printf("enter the array limit:")
	fmt.Scanf("%d", &limit)
	a := make([]int, limit)
	getarray(limit, a)
	displayarray(limit, a)
}
func getarray(limit int, a []int) {
	fmt.Printf("enter the array values:")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &a[i])
	}
}
func displayarray(limit int, a []int) {
	fmt.Printf(" array values are:")
	for i := 0; i < limit; i++ {
		fmt.Printf("%d ", a[i])
	}

}
