package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	//consider for b = 0
	if b != 0 {
		return a / b
	}
	return 0
}

func calculate(a, b int, process string) int {
	var result int
	if process == "add" {
		result = add(a, b)
	}
	if process == "sub" {
		result = subtract(a, b)
	}
	if process == "mul" {
		result = multiply(a, b)
	}
	if process == "div" {
		result = divide(a, b)
	}
	fmt.Println("The calculated value is", result)
	return result
}

func main() {
	var a, b int
	var process string
	fmt.Println("Enter the first integer from 0 to 9: ")
	fmt.Scanln(&a)
	fmt.Println("Enter process: (add, sub, mul, div)")
	fmt.Scanln(&process)
	fmt.Println("Enter the second integer from 0 to 9: ")
	fmt.Scanln(&b)

	calculate(a, b, process)
}
