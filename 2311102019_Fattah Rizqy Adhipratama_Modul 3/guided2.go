package main

import (
	"fmt"
)

func main() {
	var panjang, lebar, tinggi float64
	// Input panjang, lebar, dan tinggi balok
	fmt.Print("Masukkan panjang balok: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scanln(&lebar)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scanln(&tinggi)
	// Menghitung luas permukaan balok
	luas := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	// Menghitung volume balok
	volume := panjang * lebar * tinggi
	// Menampilkan hasil
	fmt.Printf("Luas permukaan balok adalah: %.2f\n", luas)
	fmt.Printf("Volume balok adalah: %.2f\n", volume)
}