package main

import (
	"fmt"
	"math"
)

func jarak(cx, cy, x, y int) float64 {
	return math.Sqrt(float64((cx-x)*(cx-x) + (cy-y)*(cy-y)))
}

func diDalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2 int
	var x, y int

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1):")
	fmt.Scanf("%d %d %d", &cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2):")
	fmt.Scanf("%d %d %d", &cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scanf("%d %d", &x, &y)

	diDalamLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diDalamLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
