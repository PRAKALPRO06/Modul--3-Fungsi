package main

import (
	"fmt"
)

func main() {
	var panjang float64
	var lebar float64
	var tinggi float64
	fmt.Print("Panjang Balok: ")
	fmt.Scanln(&panjang)
	fmt.Print("Lebar Balok: ")
	fmt.Scanln(&lebar)
	fmt.Print("Tinggi Balok: ")
	fmt.Scanln(&tinggi)

	luasBalok := (2 * (panjang * lebar)) + (2 * (lebar * tinggi)) + (2 * (panjang * tinggi))
	volumeBalok := panjang * lebar * tinggi
	
	fmt.Printf("Balok memiliki luas kulit %.4f cm\u00B2 dan volume %.4f cm\u00B3\n", luasBalok, volumeBalok)
}
