package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func fogoh(x int) int {
	return f(g(h(x)))
}
func gohof(x int) int {
	return g(h(f(x)))
}
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a_217, b_217, c_217 int
	// Input nilai a, b, c
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_217)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b_217)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c_217)
	// Output hasil perhitungan fungsi
	fmt.Printf("fogoh(%d) = %d\n", a_217, fogoh(a_217))
	fmt.Printf("gohof(%d) = %d\n", b_217, gohof(b_217))
	fmt.Printf("hofog(%d) = %d\n", c_217, hofog(c_217))
}
