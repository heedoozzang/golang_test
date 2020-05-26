package main

import "fmt"

func main() {
	fmt.Println("Hello, snowdeer")

	var num1 int
	fmt.Scanln(&num1)

	num2 := num1 + 1
	fmt.Printf("%d + 1 = %d", num1, num2)
}
