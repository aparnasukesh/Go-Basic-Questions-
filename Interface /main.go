package main

import "fmt"

type sample interface {
	area() int
}

type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.height * r.width
}
func main() {

	var s sample
	s = rectangle{
		width:  10,
		height: 20,
	}
	fmt.Println(s.area())

	samplePrint(10)
	samplePrint(12.5)
	samplePrint("Aparna")

}

// empty interface
func samplePrint(val interface{}) {
	fmt.Printf("%v , %T\n", val, val)
}
