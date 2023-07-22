package main

import "fmt"

func main() {
	var arr [5]int = [5]int{10, 43, 23, 30, 55}
	var i, temp, j int
	for i = 0; i < 5; i++ {
		for j = i + 1; j < 5; j++ {
			if arr[i] <= arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Printf("After sorting array :")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
