package main

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	koordinatX, koordinatY, jariJari float64
}

func titikDiDalam(lingkaran Lingkaran, x, y float64) bool {
	jarak := math.Sqrt(math.Pow(x-lingkaran.koordinatX, 2) + math.Pow(y-lingkaran.koordinatY, 2))
	return jarak <= lingkaran.jariJari
}

func tentukanPosisi(lingkaran1, lingkaran2 Lingkaran, x, y float64) string {
	diLingkaran1 := titikDiDalam(lingkaran1, x, y)
	diLingkaran2 := titikDiDalam(lingkaran2, x, y)

	if diLingkaran1 && diLingkaran2 {
		return "Titik di dalam kedua lingkaran"
	} else if diLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if diLingkaran2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar kedua lingkaran"
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var x, y float64

	fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 1 (x y r):")
	fmt.Scan(&lingkaran1.koordinatX, &lingkaran1.koordinatY, &lingkaran1.jariJari)

	fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 2 (x y r):")
	fmt.Scan(&lingkaran2.koordinatX, &lingkaran2.koordinatY, &lingkaran2.jariJari)

	fmt.Println("Masukkan koordinat titik yang akan dicek (x y):")
	fmt.Scan(&x, &y)

	hasil := tentukanPosisi(lingkaran1, lingkaran2, x, y)
	fmt.Println(hasil)
}
