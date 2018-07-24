package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var val float64
	var from, to string
	//setting flags
	flag.StringVar(&from, "from", "meters", "Unit to convert from - 'meters','feet' or 'inches' ")
	flag.StringVar(&to, "to", "feet", "Unit to convert to - 'meters','feet' or 'inches'")
	//reading flags
	flag.Parse()
	//read passed value
	if len(os.Args) == 4 {
		input, _ := strconv.ParseFloat(os.Args[3], 64)
		val = input
	} else {
		fmt.Println("Parmeterss missing plese refer help")
		return
	}
	// validate flag
	if !isFlagValid(from) {
		fmt.Println("'from' flag is invalid. Allowed flag are 'meters','feet' or 'inches'")
		return
	}
	if !isFlagValid(to) {
		fmt.Println("'to' flag is invalid. Allowed flag are 'meters','feet' or 'inches'")
		return
	}
	//calling convert
	result := converter(from, to, val)
	fmt.Println("Length", val, from, "is equivalent to", result, to)
}

func isFlagValid(flag string) (b bool) {
	switch flag {
	case "meters", "feet", "inches":
		return true
	}
	return false
}

func converter(from string, to string, val float64) (result float64) {
	switch {
	case from == "meters" && to == "inches":
		return val * 39.3701
	case from == "meters" && to == "feet":
		return val * 3.2808
	case from == "inches" && to == "meters":
		return val * 0.0254
	case from == "inches" && to == "feet":
		return val * 0.0833
	case from == "feet" && to == "inches":
		return val * 12.0
	case from == "feet" && to == "meters":
		return val * 0.3048
	}
	return 0
}
