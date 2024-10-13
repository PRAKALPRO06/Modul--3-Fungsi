package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(float64((a-c)*(a-c) + (b-d)*(b-d)))
}

func diDalam(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}


func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Print("Masukkan cx1, cy1, r1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan cx2, cy2, r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat x dan y: ")
	fmt.Scan(&x, &y)

	diLingkaran1_2311102013 := diDalam(cx1, cy1, r1, x, y)
	diLingkaran2_2311102013 := diDalam(cx2, cy2, r2, x, y)

	if diLingkaran1_2311102013 && diLingkaran2_2311102013 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1_2311102013 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2_2311102013 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}