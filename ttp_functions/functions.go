package main

import "fmt"

func main() {
	s, d := getMessage()
	fmt.Println(s, d)
}

func getMessage() (s int, d string) {
	d = "mandar"
	return 0, d
}
