package main

import (
	"fmt"
	"math"
)

func hitungJarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func dalamLingkaran(cx, cy, r, x, y float64) bool {
	distance := hitungJarak(cx, cy, x, y)
	return distance < r
}

func main() {
	var cx1, cy1, r1 float64
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 float64
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	dalamLingkaran1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
