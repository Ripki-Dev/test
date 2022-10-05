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
