package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func isInsideCircle(cx, cy, r, x, y float64) bool {
	return distance(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&x, &y)

	insideCircle1 := isInsideCircle(cx1, cy1, r1, x, y)
	insideCircle2 := isInsideCircle(cx2, cy2, r2, x, y)

	if insideCircle1 && insideCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if insideCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if insideCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
