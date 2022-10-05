package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}
func Factorial(numb int) int {
	result := 1
	for i := 1; i <= numb; i++ {
		result *= i
	}
	return result
}
