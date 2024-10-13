package main

import "fmt"

func hitungLuas(panjang, lebar, tinggi float64) float64 {
	luasPermukaan := 2 * ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi))
	return luasPermukaan
}
func hitungVolume(panjang, lebar, tinggi float64) float64 {
	volume := panjang * lebar * tinggi
	return volume
}

func main() {
	var panjang, lebar, tinggi float64
	fmt.Print("masukkan Panjang Balok : ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan Lebar Balok : ")
	fmt.Scanln(&lebar)
	fmt.Print("Masukkan Tinggi Balok : ")
	fmt.Scanln(&tinggi)

	luasPermukaan := hitungLuas(panjang, lebar, tinggi)
	volume := hitungVolume(panjang, lebar, tinggi)

	fmt.Printf("Luas permukaan balok : %.2f\n", luasPermukaan)
	fmt.Printf("Volume balok : %.2f\n", volume)
}
