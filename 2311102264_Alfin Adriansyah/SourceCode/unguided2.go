package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	diLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

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
