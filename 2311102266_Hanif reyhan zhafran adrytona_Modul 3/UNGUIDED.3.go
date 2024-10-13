package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) float64 {
	return math.Sqrt(math.Pow(float64(a-c), 2) + math.Pow(float64(b-d), 2))
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Println("Masukkan data lingkaran 1 (format: x y radius):")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Println("Masukkan data lingkaran 2 (format: x y radius):")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik (format: x y):")
	fmt.Scan(&x, &y)

	d1 := jarak(x, y, cx1, cy1)
	d2 := jarak(x, y, cx2, cy2)

	if d1 <= float64(r1) && d2 <= float64(r2) {
		fmt.Println("Titik di bawah lingkaran 1 dan 2")
	} else if d1 <= float64(r1) {
		fmt.Println("Titik di bawah lingkaran 1")
	} else if d2 <= float64(r2) {
		fmt.Println("Titik di bawah lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

// 2311102266_Hanif reyhan
