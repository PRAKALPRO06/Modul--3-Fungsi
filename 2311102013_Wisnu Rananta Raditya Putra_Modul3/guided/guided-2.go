package main

import (
    "fmt"
)

func main() {
    var panjang, lebar, tinggi float64

    fmt.Print("Masukkan panjang balok: ")
    fmt.Scanln(&panjang)
    fmt.Print("Masukkan lebar balok: ")
    fmt.Scanln(&lebar)
    fmt.Print("Masukkan tinggi balok: ")
    fmt.Scanln(&tinggi)

    luas := 2 * (panjang*lebar + panjang*tinggi + lebar*tinggi)

    volume := panjang * lebar * tinggi

    fmt.Printf("Luas permukaan balok adalah: %.2f\n", luas)
    fmt.Printf("Volume balok adalah: %.2f\n", volume)
}