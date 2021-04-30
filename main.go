package main

import (
	"calc/calculator"
	"fmt"
)

func main() {
	fmt.Println("start...")

	x := 10
	y := 5

	result := calculator.Devide(x, y)
	fmt.Println(result)

	fmt.Println("end...")
}
