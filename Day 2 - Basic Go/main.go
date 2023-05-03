package main

import (
	"fmt"
)

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var sisiAtas, sisiBawah, tinggi float64
	var luas float64

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&n)
	if isPrima(n) {
		fmt.Println(n, "adalah bilangan prima")
	} else {
		fmt.Println(n, "bukan bilangan prima")
	}

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&n)

	if n%7 == 0 {
		fmt.Println(n, "adalah bilangan kelipatan 7")
	} else {
		fmt.Println(n, "bukan bilangan kelipatan 7")
	}

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&sisiBawah)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	luas = 0.5 * (sisiAtas + sisiBawah) * tinggi

	fmt.Println("Luas trapesium adalah:", luas)
}
