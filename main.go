package main

import (
	"fmt"
)

func main() {
	//Quiz Selasa 4 Okt
	var jawaban string
	var condition = true
	role := []string{"penyihir", "guard", "warewolf"}
	fmt.Println("Pilih peran : Penyihir, Guard & Warewolf")
	for condition {
		fmt.Println("Masukkan Nama :")
		fmt.Scanln(&jawaban)
		fmt.Println(WareWolfCheck(jawaban, role))
		condition = false
	}

	//Quiz Rabu 5 Okt
	fmt.Println(Factorial(5))

	//Quiz Kamis 6 Okt
	fizzbuzz(15)
}

func fizzbuzz(n int) {
	word1 := "Fizz"
	word2 := "Buzz"
	word3 := "FizzBuzz"
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, word3)
		} else if i%3 == 0 {
			fmt.Println(i, word1)
		} else if i%5 == 0 {
			fmt.Println(i, word2)
		} else {
			fmt.Println(i)
		}
	}
}

func Factorial(numb int) int {
	result := 1
	for i := 1; i <= numb; i++ {
		result *= i
	}
	return result
}

func WareWolfCheck(usr string, role []string) string {

	for _, i2 := range role {
		if i2 == usr {
			return "Selamat Bermain"
		} else if usr == "" || usr == " " {
			return "Nama Harus Diisi"
		}
	}
	return "Pilih Peranmu Untuk Mulai"
}
