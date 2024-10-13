//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import "fmt"

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutation(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func combination(n, r int) int {
	if r > n {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// Meminta input dari pengguna dalam satu baris
	fmt.Println("Masukkan 4 angka (a b c d) yang dipisahkan oleh spasi:")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Baris pertama: hasil permutasi dan kombinasi untuk (a, c)
	P_a_c := permutation(a, c)
	C_a_c := combination(a, c)
	fmt.Println(P_a_c, C_a_c)

	// Baris kedua: hasil permutasi dan kombinasi untuk (b, d)
	P_b_d := permutation(b, d)
	C_b_d := combination(b, d)
	fmt.Println(P_b_d, C_b_d)
}
