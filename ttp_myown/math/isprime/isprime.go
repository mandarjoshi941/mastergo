package main

import (
	"fmt"
	"time"
)

func main() {
	//num := 941
	start := time.Now()
	count := 0
	for idx := 1; idx <= 941; idx++ {
		if isPrime(idx) {
			fmt.Println(idx)
			count++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time Taken:", elapsed)
	fmt.Println("Total count:", count)
}

func isPrime(num int) bool {
	if num <= 0 || num == 1 {
		return false
	}
	for idx := 2; idx < num/2; idx++ {
		if num%idx == 0 {
			return false
		}
	}
	return true
}
