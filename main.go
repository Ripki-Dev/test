package main

import "fmt"

func main() {
	//masukkan angka
	fmt.Println(Factorial(5))
}
func Factorial(numb int) int {
	result := 1
	for i := 1; i <= numb; i++ {
		result *= i
	}
	return result
}
