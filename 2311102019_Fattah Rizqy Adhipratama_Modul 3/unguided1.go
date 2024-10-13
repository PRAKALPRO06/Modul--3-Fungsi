package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Fungsi untuk menghitung permutasi (P)
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi (C)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Input data
	var a, b, c, d int
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Permutasi dan kombinasi untuk a terhadap c
	P_a_c := permutation(a, c)
	C_a_c := combination(a, c)

	// Permutasi dan kombinasi untuk b terhadap d
	P_b_d := permutation(b, d)
	C_b_d := combination(b, d)

	// Output
	fmt.Println("Permutasi dan Kombinasi untuk a terhadap c:")
	fmt.Printf("P(%d,%d) = %d\n", a, c, P_a_c)
	fmt.Printf("C(%d,%d) = %d\n", a, c, C_a_c)

	fmt.Println("Permutasi dan Kombinasi untuk b terhadap d:")
	fmt.Printf("P(%d,%d) = %d\n", b, d, P_b_d)
	fmt.Printf("C(%d,%d) = %d\n", b, d, C_b_d)
}