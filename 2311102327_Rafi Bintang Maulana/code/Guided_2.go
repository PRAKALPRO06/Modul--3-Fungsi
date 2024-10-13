package main

import "fmt"

func main() {
	var p, l, t float32

	fmt.Print("Panjang: ")
	fmt.Scanln(&p)
	fmt.Print("Lebar: ")
	fmt.Scanln(&l)
	fmt.Print("Tinggi: ")
	fmt.Scanln(&t)

	fmt.Println("\nLuas Balok: ")
	fmt.Println(luasbalok(p, l, t))
	fmt.Println("Volume Balok: ")
	fmt.Println(volume(p, l, t))
}

func luasbalok(p, l, t float32) float32 {
	return 2 * (p*l + p*t + l*t)
}

func volume(p, l, t float32) float32 {
	return p * l * t
}