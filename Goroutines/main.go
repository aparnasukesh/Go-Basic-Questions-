package main

import (
	"fmt"
	"time"
)

func main() {
	go number()
	time.Sleep(time.Millisecond)
	go name()
	time.Sleep(time.Millisecond)

}
func number() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func name() {
	for i := 0; i < 5; i++ {
		fmt.Println("Heii")
	}
}
