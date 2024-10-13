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

	fmt.Println("Masukkan bilangan bulat : ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Menghitung hasil fungsi komposisi
	fmt.Printf("fogoh(%d) = %d\n", a, fogoh(a))
	fmt.Printf("gohof(%d) = %d\n", b, gohof(b))
	fmt.Printf("hofog(%d) = %d\n", c, hofog(c))
}
