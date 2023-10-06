package main

import "fmt"

func main() {
	a := 0
	b := 1

	fmt.Printf("%d, %d, ", a, b) // Print the first two numbers

	for i := 3; i <= 15; i++ {
		c := a + b
		fmt.Print(c, ", ")
		a = b
		b = c
	}
	fmt.Println()
}
