package main

import "fmt"

// Program yang memuat fungsi faktorial dan permutasi

func faktorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}
func main() {
	var (
		a, b, c, d int
	)
	fmt.Scanln(&a, &b, &c, &d)
	if a < c || b < d {
		fmt.Println("Input tidak valid!")
		return
	}
	fmt.Printf("%d %d \n%d %d", permutasi(a, c), kombinasi(a, c), permutasi(b, d), kombinasi(b, d))
}
