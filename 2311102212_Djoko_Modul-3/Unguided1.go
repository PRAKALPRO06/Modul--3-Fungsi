package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasian(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a1, b1, a2, b2 int

	// Input nilai a1, b1, a2, b2
	fmt.Print("Masukkan nilai a1: ")
	fmt.Scan(&a1)
	fmt.Print("Masukkan nilai b1: ")
	fmt.Scan(&b1)
	fmt.Print("Masukkan nilai a2: ")
	fmt.Scan(&a2)
	fmt.Print("Masukkan nilai b2: ")
	fmt.Scan(&b2)

	// Menghitung permutasi dan kombinasi
	p1 := permutasian(a1, a2)
	p2 := permutasian(b1, b2)
	c1 := kombinasi(a1, a2)
	c2 := kombinasi(b1, b2)

	// Output hasil perhitungan
	fmt.Printf("Permutasi(%d, %d) = %d\n", a1, a2, p1)
	fmt.Printf("Kombinasi(%d, %d) = %d\n", a1, a2, c1)
	fmt.Printf("Permutasi(%d, %d) = %d\n", b1, b2, p2)
	fmt.Printf("Kombinasi(%d, %d) = %d\n", b1, b2, c2)
}
