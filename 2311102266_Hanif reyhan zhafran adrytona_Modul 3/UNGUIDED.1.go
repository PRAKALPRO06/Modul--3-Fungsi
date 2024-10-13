package main

import (
	"fmt"
)

func main() {
	var a_2311102266, b_2311102266, c_2311102266 int
	fmt.Print("Masukkan Nilai a, b, c: ")
	fmt.Scanf("%d %d %d", &a_2311102266, &b_2311102266, &c_2311102266)
	fogohResult := fogoh(a_2311102266)
	gohofResult := gohof(b_2311102266)
	hofogResult := hofog(c_2311102266)
	fmt.Printf("%d\n", fogohResult)
	fmt.Printf("%d\n", gohofResult)
	fmt.Printf("%d\n\n", hofogResult)
}

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

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 2311102266_Hanif reyhan
