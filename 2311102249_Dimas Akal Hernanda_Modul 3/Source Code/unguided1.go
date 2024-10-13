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

// Fungsi untuk menghitung kombinasi
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	fmt.Printf("Permutasi (%dP%d): %d, Kombinasi (%dC%d): %d\n", a, c, p1, a, c, c1)

	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)
	fmt.Printf("Permutasi (%dP%d): %d, Kombinasi (%dC%d): %d\n", b, d, p2, b, d, c2)
}
