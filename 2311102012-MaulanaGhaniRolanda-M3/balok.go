//2311102012
package main

import "fmt"

func main() {
	// Deklarasi variabel panjang, lebar, dan tinggi
	var panjang, lebar, tinggi float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan panjang balok: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scan(&lebar)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scan(&tinggi)

	// Menghitung luas permukaan balok
	luasPermukaan := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)

	// Menghitung volume balok
	volume := panjang * lebar * tinggi

	// Menampilkan hasil
	fmt.Printf("Luas Permukaan Balok = %.2f\n", luasPermukaan)
	fmt.Printf("Volume Balok = %.2f\n", volume)
}
