package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func diDalam(cx, cy int, r float64, x, y int) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1 int
	var cx2, cy2 int
	var r1, r2 float64
	var x, y int

	fmt.Print("Masukkan cx1, cy1, r1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan cx2, cy2, r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat x dan y: ")
	fmt.Scan(&x, &y)

	if diDalam(cx1, cy1, r1, x, y) && diDalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam(cx1, cy1, r1, x, y) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam(cx2, cy2, r2, x, y) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
