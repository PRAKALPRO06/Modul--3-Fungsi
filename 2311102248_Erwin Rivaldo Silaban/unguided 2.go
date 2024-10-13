package main

import "fmt"

func main() {
	var win_248, b, c int

	fmt.Scan(&win_248, &b, &c)

	fmt.Println(fogoh(win_248))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}

func y(x int) int {
	return x * x
}

func z(x int) int {
	return x - 2
}

func o(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return y(z(o(x)))
}

func gohof(x int) int {
	return z(o(y(x)))
}

func hofog(x int) int {
	return o(y(z(x)))
}
