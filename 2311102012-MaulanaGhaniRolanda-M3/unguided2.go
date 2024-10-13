// 2311102012
package main

import (
	"fmt"
)

// Fungsi f(x), g(x), dan h(x)
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Komposisi fungsi: f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Komposisi fungsi: g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Komposisi fungsi: h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

// Fungsi untuk menghitung faktorial
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

func main() {
	var a, b, c int
	// Membaca input
	fmt.Print("Masukkan a, b, c: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	// Menghitung hasil
	fogohResult := fogoh(a)
	gohofResult := gohof(b)
	hofogResult := hofog(c)

	// Mencetak hasil sesuai dengan format contoh
	fmt.Printf("%d\n", fogohResult)
	fmt.Printf("%d\n", gohofResult)
	fmt.Printf("%d\n\n", hofogResult)

}
