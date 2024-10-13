//2311102037 BRIAN FARREL EVANDHIKA IF 11 06
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran
func diDalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int // Lingkaran 1
	var cx2, cy2, r2 int // Lingkaran 2
	var x, y int         // Titik yang akan dicek

	// Input data lingkaran dan titik
	fmt.Print("Masukkan cx1, cy1, r1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan cx2, cy2, r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat x dan y: ")
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap dua lingkaran
	diLingkaran1 := diDalam(cx1, cy1, r1, x, y)
	diLingkaran2 := diDalam(cx2, cy2, r2, x, y)

	// Menentukan output berdasarkan kondisi
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
