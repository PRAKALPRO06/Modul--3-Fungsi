package main

import "fmt"

// Program untuk menghitung volume dan luas permukaan balok
// keterangan: a: panjang, b: lebar, c: tinggi

func lp_balok(a, b, c int) int {
	result := 2 * (a*b + a*c + b*c)
	return result
}
func v_balok(a, b, c int) int {
	return a * b * c
}

func main() {
	var a, b, c int
	fmt.Print("Masukkan panjang, lebar dan tinggi balok: ")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Volume balok adalah: ", v_balok(a, b, c))
	fmt.Println("Luas permukaan balok adalah: ", lp_balok(a, b, c))
}
