package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func diDalamLingkaran(cx, cy, r, x, y float64) bool {

	jarakKeTitik := jarak(cx, cy, x, y)

	return jarakKeTitik <= r
}

func main() {

	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var px, py float64

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (x1, y1, r1): ")
	fmt.Scan(&x1, &y1, &r1)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (x2, y2, r2): ")
	fmt.Scan(&x2, &y2, &r2)

	fmt.Print("Masukkan koordinat titik sembarang (px, py): ")
	fmt.Scan(&px, &py)

	diLingkaran1 := diDalamLingkaran(x1, y1, r1, px, py)
	diLingkaran2 := diDalamLingkaran(x2, y2, r2, px, py)

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
