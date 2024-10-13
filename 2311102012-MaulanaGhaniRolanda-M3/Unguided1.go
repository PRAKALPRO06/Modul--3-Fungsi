//2311102012
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Deklarasi variabel untuk input
	var a, b, c, d int

	// Meminta input dari pengguna dengan pengecekan syarat a >= c dan b >= d
	for {
		fmt.Println("Masukkan nilai a, b, c, dan d (syarat: a >= c dan b >= d):")
		fmt.Scan(&a, &b, &c, &d)

		// Cek apakah syarat terpenuhi
		if a >= c && b >= d {
			break
		} else {
			fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		}
	}

	// Menghitung permutasi dan kombinasi untuk (a, c) dan (b, d)
	p1 := permutation(a, c)
	c1 := combination(a, c)
	p2 := permutation(b, d)
	c2 := combination(b, d)

	// Menampilkan hasil perhitungan
	fmt.Printf("P(%d,%d) = %d\n", a, c, p1)
	fmt.Printf("C(%d,%d) = %d\n", a, c, c1)
	fmt.Printf("P(%d,%d) = %d\n", b, d, p2)
	fmt.Printf("C(%d,%d) = %d\n", b, d, c2)
}
