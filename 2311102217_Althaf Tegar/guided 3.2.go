package main

import (
	"fmt"
)

func main() {
	var panjang, lebar, tinggi_217 float64

	fmt.Print("Masukkan panjang balok: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scanln(&lebar)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scanln(&tinggi_217)

	luasPermukaan := 2 * (panjang*lebar + panjang*tinggi_217 + lebar*tinggi_217)

	volume := panjang * lebar * tinggi_217

	fmt.Printf("Luas Permukaan Balok: %.0f\n", luasPermukaan)
	fmt.Printf("Volume Balok: %.0f\n", volume)
}
