package main

import (
	"fmt"
)

func volumeBalok(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}


func luasPermukaanBalok(panjang, lebar, tinggi float64) float64 {
	return 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
}

func main() {
	
	var panjang, lebar, tinggi float64
	
	fmt.Print("Masukkan panjang balok: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scan(&tinggi)

	
	volume := volumeBalok(panjang, lebar, tinggi)
	luasPermukaan := luasPermukaanBalok(panjang, lebar, tinggi)

	// Menampilkan hasil
	fmt.Printf("Volume balok: %.2f\n", volume)
	fmt.Printf("Luas permukaan balok: %.2f\n", luasPermukaan)
}
