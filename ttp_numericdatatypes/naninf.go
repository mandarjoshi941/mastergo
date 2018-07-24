package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0.0
	fmt.Println(1 / n)         // +Inf
	fmt.Println(-1 / n)        // -Inf
	fmt.Println(n / n)         // NaN
	fmt.Println(math.Sqrt(-1)) // NaN

	if math.IsInf((1 / n), 1) { // we have to pass the expected sign
		fmt.Println("This is positive infinity")
	}

	if math.IsInf((-1 / n), -1) { // we have to pass the expected sign
		fmt.Println("This is nigative infinity")
	}

	if math.IsNaN(n / n) {
		fmt.Println("This is not a number")
	}
}
