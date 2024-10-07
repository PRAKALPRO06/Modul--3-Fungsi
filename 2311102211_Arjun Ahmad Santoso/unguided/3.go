package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1, y1, r1, x2, y2, r2, x, y int
	)
	// Keterangan:
	// (x1, y1) dan r1 berturut-turut adalah koordinat titik pusat dan jari-jari lingkaran 1
	// (x2, y2) dan r2 berturut-turut adalah koordinat titik pusat dan jari-jari lingkaran 2
	// (x, y) adalah koordinat titik yang akan dicek

	fmt.Scanln(&x1, &y1, &r1)
	fmt.Scanln(&x2, &y2, &r2)
	fmt.Scanln(&x, &y)

	s_t_in_c1 := math.Sqrt(math.Pow(float64(x1-x), 2)+math.Pow(float64(y1-y), 2)) < float64(r1) // Nilai kebenaran titik (x, y) berada di dalam lingkaran 1
	s_t_in_c2 := math.Sqrt(math.Pow(float64(x2-x), 2)+math.Pow(float64(y2-y), 2)) < float64(r2) // Nilai kebenaran titik (x, y) berada di dalam lingkaran 2

	var message string
	if s_t_in_c1 && s_t_in_c2 {
		message = "Titik di dalam lingkaran 1 dan 2"
	} else if s_t_in_c1 {
		message = "Titik di dalam lingkaran 1"
	} else if s_t_in_c2 {
		message = "Titik di dalam lingkaran 2"
	} else {
		message = "Titik di luar lingkaran 1 dan 2"
	}
	fmt.Println(message)
}
