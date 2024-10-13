package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
    hasil := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	return hasil
}

func dalamLingkaran(x, y, cx, cy, r float64) bool {
    var hasil bool
    jarak := jarak(x, y, cx, cy) 
    if jarak <= r {
        hasil = true
    } else {
        hasil = false
    }
	return hasil
}

func main() {

    var a, b, c, d, e, f float64
	var x, y float64

	fmt.Print("Masukkan koordinat pusat lingkaran 1 dan radius : ")
	fmt.Scan(&a, &b, &c)

	fmt.Print("Masukkan koordinat pusat lingkaran 2 dan radius : ")
	fmt.Scan(&d, &e, &f)

	fmt.Print("Masukkan koordinat titik sembarang : ")
	fmt.Scan(&x, &y)

	dalamLingkaran1 := dalamLingkaran(x, y, a, b, c)
	dalamLingkaran2 := dalamLingkaran(x, y, d, e, f)

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
