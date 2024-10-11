package main

import (
	"fmt"
)

// Fungsi fogoh
func fogoh(a int) int {
	h := a + 1
	g := h - 2
	f := g * g
	return f
}

// Fungsi gohof
func gohof(a int) int {
	f := a * a
	h := f + 1
	g := h - 2
	return g
}

// Fungsi hofog
func hofog(a int) int {
	g := a - 2
	f := g * g
	h := f + 1
	return h
}

func main() {
	var a, b, c int

	// Input nilai a, b, c
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	// Output hasil perhitungan fungsi
	fmt.Printf("fogoh(%d) = %d\n", a, fogoh(a))
	fmt.Printf("gohof(%d) = %d\n", b, gohof(b))
	fmt.Printf("hofog(%d) = %d\n", c, hofog(c))
}
