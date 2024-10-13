package main

import "fmt"

var p, l, t int

func luasPermukaanBalok(p, l, t int) int {
	luas := 2 * (p*l + l*t + p*t)
	return luas
}

func volumeBalok(p, l, t int) int {
	volume := p * l * t
	return volume
}

func input() {
	fmt.Print("Masukkan panjang balok : ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar balok : ")
	fmt.Scan(&l)
	fmt.Print("Masukkan tinggi balok : ")
	fmt.Scan(&t)
}

func main() {

	input()

	lausBalok := luasPermukaanBalok(p, l, t)
	volume := volumeBalok(p, l, t)

	fmt.Printf("Luas permukaan baloknya adalah %v dan volumenya adalah %v", lausBalok, volume)
}
