package main

import "fmt"

func f(x_226 int) int {
	return x_226 * x_226
}

func g(x_226 int) int {
	return x_226 - 2
}

func h(x_226 int) int {
	return x_226 + 1
}

func main() {
	var a_226, b_226, c_226 int
	fmt.Printf("Masukkan tiga bilangan (dipisahkan oleh spasi): ")
	fmt.Scan(&a_226, &b_226, &c_226)

	fogoh := f(g(h(a_226)))
	gohof := g(h(f(b_226)))
	hofog := h(f(g(c_226)))

	fmt.Println(fogoh)
	fmt.Println(gohof)
	fmt.Println(hofog)
}
