package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	if len(os.Args) > 1 {
		input, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		n = input
	} else {
		fmt.Println("Plese enter a number grater than 1:")
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	playfizzbuzz(n)
}

func playfizzbuzz(num int) {
	for idx := 1; idx <= num; idx++ {
		switch {
		case idx%15 == 0:
			fmt.Println("FizzBuzz")
		case idx%3 == 0:
			fmt.Println("Fizz")
		case idx%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(idx)
		}
	}
}
