package main

import (
	"fmt"
	"math/big"
)

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		result.Mul(result, big.NewInt(i))
	}
	return result
}

func permutation(n, r int64) *big.Int {
	return new(big.Int).Div(factorial(n), factorial(n-r))
}

func combination(n, r int64) *big.Int {
	return new(big.Int).Div(factorial(n), new(big.Int).Mul(factorial(r), factorial(n-r)))
}

func main() {
	var a, b, c, d int64
	fmt.Print("Masukkan empat bilangan (a, b, c, d) dipisahkan oleh spasi: ")
	_, err := fmt.Scan(&a, &b, &c, &d)

	if err != nil {
		fmt.Println("Input tidak valid. Silakan masukkan empat bilangan.")
		return
	}

	if a < c || b < d {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		return
	}

	fmt.Printf("P(%d, %d) = %v\n", a, c, permutation(a, c))
	fmt.Printf("C(%d, %d) = %v\n", a, c, combination(a, c))
	fmt.Printf("P(%d, %d) = %v\n", b, d, permutation(b, d))
	fmt.Printf("C(%d, %d) = %v\n", b, d, combination(b, d))
}
