package main

import (
	"calc/mathmethod"
	"flag"
	"fmt"
)

func main() {

	var action = flag.String("action", "add", "input the math method")
	var v1 = flag.Float64("v1", 0, "input first number")
	var v2 = flag.Float64("v2", 0, "input second number")
	var v3 bool
	flag.BoolVar(&v3, "v3", true, "input third number")
	flag.Parse()

	switch *action {
	case "add":
		result := mathmethod.Add(*v1, *v2)
		fmt.Printf("add: %f + %f = %f", *v1, *v2, result)
		fmt.Print("\n")
	case "subtract":
		result := mathmethod.Subtract(*v1, *v2)
		fmt.Printf("subtract:%f - %f = %f", *v1, *v2, result)
		fmt.Print("\n")
	case "multiply":
		result := mathmethod.Multiply(*v1, *v2)
		fmt.Printf("multiply:%f x %f = %f", *v1, *v2, result)
		fmt.Print("\n")
	case "division":
		if *v2 == 0 {
			panic("params error")
		}

		result := mathmethod.Division(*v1, *v2)
		fmt.Printf("division:%f / %f = %f", *v1, *v2, result)
		fmt.Print("\n")
	case "sqrt":
		if *v1 < 0 {
			panic("params error")
		}

		result := mathmethod.Sqrt(*v1)
		fmt.Printf("Sqrt: %f is root = %f", *v1, result)
		fmt.Print("\n")
		fallthrough
	default:
		fmt.Println("error action")
	}

	fmt.Println("v3=", v3)
}
