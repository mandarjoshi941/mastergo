package main

import "fmt"

func trace(name string) func() {
	fmt.Println("Entering", name)
	return func() {
		fmt.Println("Leaving", name)
	}
}
func f() {
	defer trace2(trace("f"))
	fmt.Println("Doing something")
}

func trace2(f func()) {
	defer f()
	fmt.Println("Entering trace2")
}
func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
