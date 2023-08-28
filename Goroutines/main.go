package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go a(&wg)
	go b(&wg)

	wg.Wait()

}

func a(c *sync.WaitGroup) {
	defer c.Done()
	fmt.Println("a")
}

func b(c *sync.WaitGroup) {
	defer c.Done()
	fmt.Println("b")
}
