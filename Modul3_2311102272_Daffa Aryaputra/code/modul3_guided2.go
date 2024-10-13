package main

import (
	"fmt"
)

func main() {
	// Membuat reader untuk input dari user
	//reader := bufio.NewReader(os.Stdin)

	nama := "Hello World"

	// Menampilkan prompt kepada pengguna
	fmt.Println(nama)

	// Membaca input dari pengguna hingga baris baru (newline)
	//nama, _ := reader.ReadString('\n')

	// Menampilkan input yang diterima
	//fmt.Println("Halo,", nama)

	fmt.Println()
	balok()
}

func balok() {
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