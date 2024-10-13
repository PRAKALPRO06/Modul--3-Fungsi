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
	fmt.Print("masukan panjang permukaan: ")
	fmt.Scanln(&panjang)
	fmt.Print("masukan lebar permukaan: ")
	fmt.Scanln(&lebar)
	fmt.Print("masukan tinggi permukaan: ")
	fmt.Scanln(&tinggi)

	fmt.Println("Permukaan balok: ", cekLuas(panjang, lebar, tinggi))
	fmt.Println("Volume: ", volume(panjang, lebar, tinggi))
}
