// 2311102033_andika indra prastawa
package main

import (
	"fmt"
)

func hitungVolume(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

func hitungLuasPermukaan(panjang, lebar, tinggi float64) float64 {
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

	volume := hitungVolume(panjang, lebar, tinggi)
	luasPermukaan := hitungLuasPermukaan(panjang, lebar, tinggi)

	fmt.Printf("Volume balok: %.2f\n", volume)
	fmt.Printf("Luas permukaan balok: %.2f\n", luasPermukaan)
}
