package main

import "fmt"

type s1 struct {
	n int
	b bool
}

type s2 struct {
	r []rune
}

type s3 struct {
	r [3]rune
}

func main() {
	// try0
	s11 := s1{n: 4, b: true}
	s12 := s1{n: 4, b: true}
	fmt.Println(s11 == s12)
	// try1
	s21 := s2{r: []rune{'a', 'b', '🎵'}}
	s22 := s2{r: []rune{'a', 'b', '🎶'}}
	fmt.Println(s21.r[1] == s22.r[1])
	// try2
	s31 := s3{r: [3]rune{'a', 'b', '🎵'}}
	s32 := s3{r: [3]rune{'a', 'b', '🎶'}}
	fmt.Println(s31 == s32)
}
