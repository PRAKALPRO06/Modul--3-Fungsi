package main

import "fmt"

// Fungsi untuk menghitung volume balok
func volumeBalok(p, l, t float64) float64 {
	return p * l * t
}

// Fungsi untuk menghitung luas permukaan balok
func luasBalok(p, l, t float64) float64 {
	return 2 * ((p * l) + (p * t) + (l * t))
}

func main() {
	var p, l, t float64

	// Input nilai panjang (p), lebar (l), dan tinggi (t)
	fmt.Print("Masukkan panjang balok: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scan(&l)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scan(&t)

	// Output hasil volume dan luas permukaan balok
	fmt.Printf("Volume dari balok tersebut adalah: %.2f\n", volumeBalok(p, l, t))
	fmt.Printf("Luas dari balok tersebut adalah: %.2f\n", luasBalok(p, l, t))
}
