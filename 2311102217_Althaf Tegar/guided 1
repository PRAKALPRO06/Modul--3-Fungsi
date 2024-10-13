package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a dan b: ")
	fmt.Scan(&a, &b)

	// Memastikan a >= b
	if a >= b {
		fmt.Println("Hasil permutasi:", permutasi(a, b))
	} else {
		fmt.Println("Hasil permutasi:", permutasi(b, a))
	}
}

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int {
	if n < r {
		fmt.Println("Error: n harus lebih besar atau sama dengan r")
		return -1
	}
	return faktorial(n) / faktorial(n-r)
}
