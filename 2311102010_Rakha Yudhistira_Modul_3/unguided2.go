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

func main() {
	var a, b, c int
	fmt.Printf("Masukkan Nilai : ")
	fmt.Scan(&a, &b, &c)

	fgoh := f(g(h(a)))      
	gohof := g(h(f(b)))     
	hofog := h(f(g(c)))     

	fmt.Println(fgoh)
	fmt.Println(gohof)
	fmt.Println(hofog)
}
