package main

import "fmt"

func hitungLuasPermukaanBalok(panjang, lebar, tinggi float64) float64 {
	
	luasPermukaan := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	return luasPermukaan
}

func hitungvolumebalok(panjang, lebar, tinggi float64) float64 {
	volume := panjang*lebar*tinggi 
	return volume
}

func main() {
	var panjang, lebar, tinggi float64

	fmt.Print("Masukkan panjang balok: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scan(&tinggi)

	luasPermukaan := hitungLuasPermukaanBalok(panjang, lebar, tinggi)
	volume := hitungvolumebalok(panjang, lebar, tinggi)

	fmt.Printf("Luas permukaan balok: %.2f\n", luasPermukaan)
	fmt.Printf("Luas volume balok : %.2f\n", volume)  
}