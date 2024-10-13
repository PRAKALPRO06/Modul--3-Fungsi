// 2311102033_andika indra prastawa
package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var x, y int

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (x1 y1 r1):")
	fmt.Scan(&x1, &y1, &r1)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (x2 y2 r2):")
	fmt.Scan(&x2, &y2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&x, &y)

	jarakKeLingkaran1 := jarak(x1, y1, x, y)
	jarakKeLingkaran2 := jarak(x2, y2, x, y)

	dalamLingkaran1 := jarakKeLingkaran1 <= float64(r1)
	dalamLingkaran2 := jarakKeLingkaran2 <= float64(r2)

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
