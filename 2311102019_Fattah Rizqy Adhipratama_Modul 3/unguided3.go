package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (x1, y1) dan (x2, y2)
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func dalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	// Input data pusat lingkaran 1 dan 2, serta radius lingkaran 1 dan 2
	var cx1, cy1, r1, cx2, cy2, r2 float64
	fmt.Print("Masukkan pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Print("Masukkan pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input koordinat titik yang akan diperiksa
	var x, y float64
	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	// Cek apakah titik berada di dalam lingkaran 1 dan/atau lingkaran 2
	dalamL1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalamL2 := dalamLingkaran(cx2, cy2, r2, x, y)

	// Output hasil
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}