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

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	} else {
		fmt.Println(permutasi(b, a))
	}
}
