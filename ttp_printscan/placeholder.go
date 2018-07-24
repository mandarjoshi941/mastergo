package main

import (
	"fmt"
)

// Car is for use
type Car struct {
	Name  string
	Model string
}

func main() {
	print()
	print1()
	scan()
}

func print() {
	intVar := 1234
	stringVar := "Mandar"
	boolVar := true
	floatVar := 22.0 / 7.0

	car := Car{"i10", "Grand"}

	fmt.Printf("intVar = %d\n", intVar)
	fmt.Printf("stringVar = %s\n", stringVar)
	fmt.Printf("boolVar = %t\n", boolVar)
	fmt.Printf("floatVar = %f\n", floatVar)
	fmt.Printf("Anything = %v\n", car)
	fmt.Printf("Anything = %#v\n", car)
	fmt.Printf("Anything = %T\n", car)
	fmt.Printf("%[2]f %[1]X\n", intVar, floatVar)
}

func scan() {
	// var a float64
	// fmt.Println("Please enter a float value")
	// fmt.Scanf("%f", &a)
	// fmt.Println("The float you entered is :", a)

	var n1, n2, n3, n4 int
	var f1 float64
	// Scan the card number.
	str1 := "Card number: 1234 5678 0123 4567"
	_, err := fmt.Sscanf(str1, "Card number: %d %d %d %d", &n1, &n2, &n3, &n4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%04d %04d %04d %04d\n", n1, n2, n3, n4)
	// Scan the numeric values into a floating-point variable, and an integer.
	str2 := "Brightness is 50.0% (hex #7ffff)"
	_, err = fmt.Sscanf(str2, "Brightness is %f%% (hex #%x)", &f1, &n1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f1, n1)
}

func print1() {
	// Print RGB values...
	r, g, b := 124, 87, 3
	// ...as #7c5703  (specifying hex format, fixed width, and leading zeroes)
	// Hint: don't forget to add a newline at the end of the format string.
	fmt.Printf("#%#X #%#X #%#X \n", r, g, b)
	// ...as rgb(124, 87, 3)
	fmt.Printf("rgb(%d, %d, %d)\n", r, g, b)
	// ...as rgb(124, 087, 003) (specifying fixed width and leading zeroes)
	fmt.Printf("rgb(%03d, %03d, %03d)\n", r, g, b)
	// ...as rgb(48%, 34%, 1%) (specifying a literal percent sign)
	fmt.Printf("rgb(%d%%, %d%%, %d%%)\n", r*100/255, g*100/255, b*100/255)
	// Print the type of r.
	fmt.Printf("%T\n", r)
}
