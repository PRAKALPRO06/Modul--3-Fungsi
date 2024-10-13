package main

import "fmt"

var a, b, c, d int

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	fmt.Print("Masukkan input (a b c d) = ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Printf("Permutasi(a, c): %d, Kombinasi(a, c): %d\n", permutasi(a, c), kombinasi(a, c))

		fmt.Printf("Permutasi(b, d): %d, Kombinasi(b, d): %d\n", permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}
