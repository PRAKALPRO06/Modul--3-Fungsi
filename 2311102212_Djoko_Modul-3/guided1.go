package main

import "fmt"

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func main() {
	var a, b int

	// Input nilai a dan b
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Periksa kondisi dan hitung permutasi sesuai nilai a dan b
	if a >= b {
		fmt.Printf("Permutasi(%d, %d) = %d\n", a, b, permutasi(a, b))
	} else {
		fmt.Printf("Permutasi(%d, %d) = %d\n", b, a, permutasi(b, a))
	}
}
