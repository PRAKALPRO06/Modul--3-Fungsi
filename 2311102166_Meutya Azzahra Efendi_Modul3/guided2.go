package main

import "fmt"

func hitungLuasPermukaan(panjang, lebar, tinggi int) int {
	return 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
}
func hitungVolume(panjang, lebar, tinggi int) int {
	return panjang + lebar + tinggi
}
func main() {
	var panjang, lebar, tinggi int

	fmt.Println("Masukkan panjang balok: ")
	fmt.Scan(&panjang)
	fmt.Println("Masukkan lebar balok: ")
	fmt.Scan(&lebar)
	fmt.Println("Masukkan tinggi balok: ")
	fmt.Scan(&tinggi)

	luasPermukaan := hitungLuasPermukaan(panjang, lebar, tinggi)
	volume := hitungVolume(panjang, lebar, tinggi)

	fmt.Printf("Luas permukaan balok: %d\n", luasPermukaan)
	fmt.Printf("Volume balok: %d\n",volume)
}