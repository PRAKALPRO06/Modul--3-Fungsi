package main

import "fmt"

func hitungVolumeBalok(panjang, lebar, tinggi float64) float64 {
    return panjang * lebar * tinggi
}

func hitungLuasPermukaanBalok(panjang, lebar, tinggi float64) float64 {
    return 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
}

func main() {
    var panjang, lebar, tinggi float64

    fmt.Print("Masukkan panjang balok\t: ")
    fmt.Scanln(&panjang)
    fmt.Print("Masukkan lebar balok\t: ")
    fmt.Scanln(&lebar)
    fmt.Print("Masukkan tinggi balok\t: ")
    fmt.Scanln(&tinggi)

    luasPermukaan := hitungLuasPermukaanBalok(panjang, lebar, tinggi)
	volume := hitungVolumeBalok(panjang, lebar, tinggi)

    fmt.Printf("Luas permukaan balok adalah\t: %.2f\n", luasPermukaan)
	fmt.Printf("Volume balok adalah\t\t: %.2f\n", volume)
}
