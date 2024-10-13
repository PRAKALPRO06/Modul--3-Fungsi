package main

import "fmt"

func main() {
	var p, l, t int

	fmt.Print("masukkan panjang, lebar, dan tinggi balok:")
	fmt.Scan(&p, &l, &t)

	fmt.Println("luas permukaan  & volume adalah :", luasp(p, l, t))
	fmt.Println(hitungvolume(p, l, t))
}

func luasp(p, l, t int) int {
	return (2*p*l + p*t + l*t)

}

func hitungvolume(p, l, t int) int {
	return (p * l * t)

}
