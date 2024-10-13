package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}
func Permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(Permutasi(a, b))
	} else {
		fmt.Println(Permutasi(b, a))
	}
}
