package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	a, b = b, a
	fmt.Printf("a : %d , b : %d", a, b)
}
