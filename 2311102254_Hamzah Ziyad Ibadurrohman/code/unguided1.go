package main

import "fmt"

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
	return g(f(h(x)))
}
func hogof(x int) int {
	return h(g(f(x)))
}

func main() {
	var a, b, c int

	fmt.Print("Input nilai a, b, dan c: ")
	fmt.Scan(&a, &b, &c)
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hogof(c))

}
