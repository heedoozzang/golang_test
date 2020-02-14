package main

import "fmt"

func main() {
	fmt.Println("Hello, snowdeer")

	var num int
	fmt.Scanln(&num)

	num2 := num + 1
	fmt.Printf("%d + 1 = %d", num, num2)
}
