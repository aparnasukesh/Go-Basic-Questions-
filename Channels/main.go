package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)
	c <- 20
	c <- 30
	c <- 40
	go func() {
		time.Sleep(time.Second)
		fmt.Println(<-c)
		fmt.Println(<-c)
		fmt.Println(<-c)
	}()

	time.Sleep(2 * time.Second)
}
