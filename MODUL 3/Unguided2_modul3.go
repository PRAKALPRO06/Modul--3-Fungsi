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
	fmt.Print("Masukkan nilai a, b, c : ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println("Hasil dari fogoh(a) : ", fogoh(a))
	fmt.Println("Hasil dari gohof(b) : ", gohof(b))
	fmt.Println("Hasil dari hofog(c) : ", hofog(c))
}
