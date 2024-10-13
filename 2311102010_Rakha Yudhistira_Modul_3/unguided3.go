package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y float64
	fmt.Scan(&x1, &y1, &r1) 
	fmt.Scan(&x2, &y2, &r2) 
	fmt.Scan(&x, &y)        

	diDalam1 := diDalam(x1, y1, r1, x, y)
	diDalam2 := diDalam(x2, y2, r2, x, y)

	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam Lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam Lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam Lingkaran 2")
	} else {
		fmt.Println("Titik di luar Lingkaran 1 dan 2")
	}
}
