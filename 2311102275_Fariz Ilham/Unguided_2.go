package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a,b) dan (c,d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

// Fungsi untuk mengecek apakah suatu titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	// Input koordinat titik pusat dan radius lingkaran 1
	var cx1, cy1, r1 float64
	fmt.Println("Masukkan koordinat pusat lingkaran 1 (cx1, cy1) dan radiusnya (r1):")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input koordinat titik pusat dan radius lingkaran 2
	var cx2, cy2, r2 float64
	fmt.Println("Masukkan koordinat pusat lingkaran 2 (cx2, cy2) dan radiusnya (r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input koordinat titik sembarang
	var x, y float64
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scan(&x, &y)

	// Mengecek apakah titik berada di dalam lingkaran 1 atau 2
	lingkaran1 := didalam(cx1, cy1, r1, x, y)
	lingkaran2 := didalam(cx2, cy2, r2, x, y)

	// Output hasil berdasarkan posisi titik
	if lingkaran1 && lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
