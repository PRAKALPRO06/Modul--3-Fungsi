package main

import (
	"fmt"
)

func main() {
	balok()
}

func balok() {
	var panjang, lebar, tinggi float64
	fmt.Print("Panjang = ")
	fmt.Scanln(&panjang)
	fmt.Print("Lebar = ")
	fmt.Scanln(&lebar)
	fmt.Print("Tinggi = ")
	fmt.Scanln(&tinggi)

	
	luasPermukaan := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	volume := panjang * lebar * tinggi

	fmt.Println("Luas Permukaan = ", luasPermukaan)
	fmt.Println("Volume = ", volume)
}
