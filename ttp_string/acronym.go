package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func acronym(s string) (acr string) {
	flag := true
	for _, ch := range s {
		if flag && unicode.IsLetter(ch) {
			acr = acr + string(unicode.ToUpper(ch))
			flag = false
		}
		if unicode.IsSpace(ch) {
			flag = true
		}
	}
	return acr
}

func main() {
	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
