package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 2)
}

func g(x float64) float64 {
	return x - 2
}

func h(x float64) float64 {
	return x + 1
}

func fogoh(x float64) float64 {
	return f(g(h(x)))
}

func gohof(x float64) float64 {
	return g(h(f(x)))
}

func hofog(x float64) float64 {
	return h(f(g(x)))
}

func main() {
	var a, b, c float64
	fmt.Printf("Masukkan tiga angka (dipisah dengan spasi):")
	_, err := fmt.Scanf("%f %f %f", &a, &b, &c)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Printf("fogoh dari %.f yaitu : %.f\n", a, fogoh(a))
	fmt.Printf("gohof dari %.f yaitu : %.f\n", b, gohof(b))
	fmt.Printf("hofog dari %.f yaitu : %.f\n", c, hofog(c))
}
