package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung kuadrat
func kuadrat(a float64) float64 {
	return a * a
}

// Fungsi untuk mencari posisi titik sembarang terhadap lingkaran
func mencariTitikSembarang(cx, cy, r, x, y float64) string {
	// Menghitung jarak Euclidean antara titik (x, y) dan pusat lingkaran (cx, cy)
	jarak := math.Sqrt(kuadrat(x-cx) + kuadrat(y-cy))

	// Menentukan apakah titik di dalam, di luar, atau di tepi lingkaran
	if jarak < r {
		return "dl" // dl: di dalam lingkaran
	} else if jarak == r {
		return "tl" // tl: di tepi lingkaran
	} else {
		return "ll" // ll: di luar lingkaran
	}
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	// Input pusat dan radius lingkaran pertama
	fmt.Print("Masukkan cx1, cy1, r1 untuk lingkaran 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input pusat dan radius lingkaran kedua
	fmt.Print("Masukkan cx2, cy2, r2 untuk lingkaran 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik sembarang (x, y)
	fmt.Print("Masukkan titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	// Mencari posisi titik sembarang terhadap kedua lingkaran
	d1 := mencariTitikSembarang(cx1, cy1, r1, x, y)
	d2 := mencariTitikSembarang(cx2, cy2, r2, x, y)

	// Menentukan posisi titik relatif terhadap lingkaran 1 dan 2
	if d1 == "dl" && d2 == "dl" {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 == "ll" && d2 == "ll" {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	} else if d1 == "dl" && d2 == "ll" {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d1 == "ll" && d2 == "dl" {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di tepi lingkaran")
	}
}
