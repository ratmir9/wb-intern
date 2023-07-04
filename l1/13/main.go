package main

import "fmt"

func swapNumbers(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}

func main() {
	num1 := 15
	num2 := 30

	fmt.Printf("Исходные значения: num1 = %d, num2 = %d\n", num1, num2)

	num1, num2 = swapNumbers(num1, num2)

	fmt.Printf("После обмена: num1 = %d, num2 = %d\n", num1, num2)
}
