package main

import "fmt"

func cekLuas(panjang, lebar, tinggi float64) float64 {
	return 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
}

func volume(panjang, lebar, tinggi float64) float64 {
	return panjang * lebar * tinggi
}

func main() {
	var panjang, lebar, tinggi float64

	fmt.Print("Masukkan panjang permukaan: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar permukaan: ")
	fmt.Scanln(&lebar)
	fmt.Print("Masukkan tinggi permukaan: ")
	fmt.Scanln(&tinggi)

	fmt.Println("Luas permukaan balok: ", cekLuas(panjang, lebar, tinggi))
	fmt.Println("Volume balok: ", volume(panjang, lebar, tinggi))
}
