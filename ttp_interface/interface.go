// this is the main package
package main

import (
	"fmt"
)

// Common interface
type Common interface {
}

// Parent is interfce
type Parent interface {
	A()
}

//Child struct
type Child struct{}

// A is part of Child
func (c Child) A() {
	fmt.Println("In Child A")
}

// B is here
func (c *Child) B() {
	fmt.Println("In Child B")
}
func main() {
	var p Parent
	p = Child{}
	callA(p)
}

func callA(p Parent) {
	p.A()
}
