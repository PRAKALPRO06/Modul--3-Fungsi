package main

import "fmt"

// Program yang memuat fungsi faktorial dan permutasi

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func main() {
	var (
		a, b, c int
	)
	fmt.Scanln(&a, &b, &c)
	fmt.Printf("%d\n", f(g(h(a))))
	fmt.Printf("%d\n", g(h(f(b))))
	fmt.Printf("%d\n", h(f(g(c))))
}
