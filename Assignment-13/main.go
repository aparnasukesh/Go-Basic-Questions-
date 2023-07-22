package main

import (
	"fmt"
)

func main() {
	var name string
	flag := 0
	fmt.Printf("Enter the string : ")
	fmt.Scan(&name)
	for i := 0; i < len(name); i++ {
		if name[i] != name[len(name)-i-1] {
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("String is palindrome")
	} else {
		fmt.Println("String is not palindrome")
	}
}
