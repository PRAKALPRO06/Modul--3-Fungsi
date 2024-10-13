package main

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	cx, cy, r_226 float64
}

func inputLingkaran(urutan226 int) Lingkaran {
	var L Lingkaran
	fmt.Printf("Masukkan koordinat titik pusat lingkaran %d (x y): ", urutan226)
	fmt.Scanln(&L.cx, &L.cy)
	fmt.Printf("Masukkan radius lingkaran %d: ", urutan226)
	fmt.Scanln(&L.r_226)
	return L
}

func inputTitik() (float64, float64) {
	var x, y float64
	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scanln(&x, &y)
	return x, y
}

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func cekPosisiTitik(L1, L2 Lingkaran, x, y float64) string {
	jarak1 := hitungJarak(x, y, L1.cx, L1.cy)
	jarak2 := hitungJarak(x, y, L2.cx, L2.cy)

	if jarak1 <= L1.r_226 && jarak2 <= L2.r_226 {
		return "Titik berada di dalam lingkaran 1 dan 2"
	} else if jarak1 <= L1.r_226 {
		return "Titik berada di dalam lingkaran 1"
	} else if jarak2 <= L2.r_226 {
		return "Titik berada di dalam lingkaran 2"
	} else {
		return "Titik berada di luar lingkaran 1 dan 2"
	}
}

func main() {
	lingkaran1 := inputLingkaran(1)
	lingkaran2 := inputLingkaran(2)
	x, y := inputTitik()

	fmt.Println(cekPosisiTitik(lingkaran1, lingkaran2, x, y))
}
