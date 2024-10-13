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
	var a, b, c int
	fmt.Println("Masukkan nilai : ")
	fmt.Scan(&a, &b, &c)
	fogohA := fogoh(a)
	gohofB := gohof(b)
	hofogC := hofog(c)
	fmt.Printf("fogoh(%d) = %d\n", a, fogohA)
	fmt.Printf("gohof(%d) = %d\n", b, gohofB)
	fmt.Printf("hofog(%d) = %d\n", c, hofogC)
}
