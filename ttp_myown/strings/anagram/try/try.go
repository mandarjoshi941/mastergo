package main

import (
	"fmt"
)

func main() {
	str := "Mandar Joshi"
	fmt.Println(str)
	runeArray := []rune(str)
	copied := make([]rune, len(runeArray))
	copy(copied, runeArray)
	fmt.Println(runeArray)
	fmt.Println(copied)
}
