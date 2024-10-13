package main

import "fmt"

func luasBalok(p, l, t int)int{
	hasil := 2 *(p * l + p * t + l*t)
	return hasil
}

func volumeBalok(p, l, t int)int{
	volume := p*l*t
	return volume
}

func main (){
	var p, l, t int

	fmt.Print("Masukkan panjang balok: ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar balok: ")
	fmt.Scan(&l)
	fmt.Print("Masukkan tinggi balok: ")
	fmt.Scan(&t)

	luasPermukaan := luasBalok(p, l, t)
	volume := volumeBalok(p, l, t)

	fmt.Printf("Luas permukaan balok: %.2f\n", luasPermukaan)
	fmt.Printf("Luas volume balok : %.2f\n",volume)
}