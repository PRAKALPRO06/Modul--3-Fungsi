package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Println("Permutasi:", permutasi(a, c))
		fmt.Println("Kombinasi:", kombinasi(a, c))
		fmt.Println("Permutasi:", permutasi(b, d))
		fmt.Println("Kombinasi:", kombinasi(b, d))
	} else {
		fmt.Println("Tidak memenuhi kondisi")
	}
}
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
