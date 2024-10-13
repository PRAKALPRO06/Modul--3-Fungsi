package main

import (
	"fmt"
)

// Fungsi f(x), g(x), dan h(x)
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Fungsi komposisi untuk menerima tiga fungsi dan menerapkannya secara berurutan
func compose3(f1, f2, f3 func(int) int, x int) int {
	return f1(f2(f3(x)))
}

func main() {
	var a, b, c int
	// Membaca input
	fmt.Print("Masukkan a, b, c: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Menghitung hasil
	fogohResult := compose3(f, g, h, a)
	gohofResult := compose3(g, h, f, b)
	hofogResult := compose3(h, f, g, c)

	// Mencetak hasil
	fmt.Printf("%d\n%d\n\n%d\n", fogohResult, gohofResult, hofogResult)
}
