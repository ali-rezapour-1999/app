package main

import (
	"fmt"
)

func main() {

	var number1, number2 int

	fmt.Println("enter number 1:")
	fmt.Scanln(&number1)

	fmt.Println("enter number 2:")
	fmt.Scanln(&number2)

	sum := number1 + number2
	fmt.Println(sum)
