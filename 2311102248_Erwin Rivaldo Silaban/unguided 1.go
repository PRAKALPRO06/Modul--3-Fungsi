package main

import "fmt"

func main() {
	var nilai1, nilai2, nilai3, nilai4 int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai1, &nilai2, &nilai3, &nilai4)
	if nilai1 >= nilai3 && nilai2 >= nilai4 {
		fmt.Println("Permutasi 1:", hitungPermutasi(nilai1, nilai3))
		fmt.Println("Kombinasi 1:", hitungKombinasi(nilai1, nilai3))
		fmt.Println("Permutasi 2:", hitungPermutasi(nilai2, nilai4))
		fmt.Println("Kombinasi 2:", hitungKombinasi(nilai2, nilai4))
	} else {
		fmt.Println("Tidak memenuhi kondisi")
	}
}

func hitungFaktorial(angka int) int {
	hasil_248 := 1
	for i := 1; i <= angka; i++ {
		hasil_248 *= i
	}
	return hasil_248
}

func hitungPermutasi(n, r int) int {
	return hitungFaktorial(n) / hitungFaktorial(n-r)
}

func hitungKombinasi(n, r int) int {
	return hitungFaktorial(n) / (hitungFaktorial(r) * hitungFaktorial(n-r))
}
