package main

import (
	"fmt"
)

func main() {

	var panjang, lebar, tinggi float64

	fmt.Print("Panjang = ")
	fmt.Scanln(&panjang)

	fmt.Print("Lebar = ")
	fmt.Scanln(&lebar)

	fmt.Print("tinggi = ")
	fmt.Scanln(&tinggi)

	luaspermukaan := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)
	volume := panjang * lebar * tinggi

	fmt.Println("luaspermukaan = ", luaspermukaan)
	fmt.Println("volume = ", volume)

}
