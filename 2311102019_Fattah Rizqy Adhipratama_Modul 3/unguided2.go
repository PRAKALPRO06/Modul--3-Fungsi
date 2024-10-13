package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi fogoh(x) = f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi gohof(x) = g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi hofog(x) = h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// Input data
	var a, b, c int
	fmt.Print("Masukkan nilai a, b, c: ")
	fmt.Scan(&a, &b, &c)

	// Menghitung fogoh(a), gohof(b), hofog(c)
	fogoh_a := fogoh(a)
	gohof_b := gohof(b)
	hofog_c := hofog(c)

	// Output
	fmt.Println("Hasil komposisi fungsi:")
	fmt.Printf("fogoh(%d) = %d\n", a, fogoh_a)
	fmt.Printf("gohof(%d) = %d\n", b, gohof_b)
	fmt.Printf("hofog(%d) = %d\n", c, hofog_c)
}