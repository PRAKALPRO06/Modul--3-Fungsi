package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak kuadrat antara dua titik
func jarakKuadrat(x1, y1, x2, y2 int) float64 {
	return math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2)
}

// Fungsi untuk menentukan posisi titik terhadap lingkaran
func posisiTitik(ex1, cy1, r1, ex2, cy2, r2, x, y int) string {
	// Hitung jarak kuadrat dari titik ke pusat lingkaran 1 dan 2
	jarak1 := jarakKuadrat(ex1, cy1, x, y)
	jarak2 := jarakKuadrat(ex2, cy2, x, y)

	// Hitung radius kuadrat untuk perbandingan
	radius1Kuadrat := math.Pow(float64(r1), 2)
	radius2Kuadrat := math.Pow(float64(r2), 2)

	inLingkaran1 := jarak1 < radius1Kuadrat
	inLingkaran2 := jarak2 < radius2Kuadrat
	onLingkaran1 := jarak1 == radius1Kuadrat
	onLingkaran2 := jarak2 == radius2Kuadrat

	if inLingkaran1 && inLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if inLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if inLingkaran2 {
		return "Titik di dalam lingkaran 2"
	} else if onLingkaran1 && onLingkaran2 {
		return "Titik di tepi lingkaran 1 dan 2"
	} else if onLingkaran1 {
		return "Titik di tepi lingkaran 1"
	} else if onLingkaran2 {
		return "Titik di tepi lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var ex1, cy1, r1 int
	var ex2, cy2, r2 int
	var x_217, y_217 int

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1:")
	fmt.Scan(&ex1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2:")
	fmt.Scan(&ex2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&x_217, &y_217)

	hasil := posisiTitik(ex1, cy1, r1, ex2, cy2, r2, x_217, y_217)
	fmt.Println(hasil)
}
