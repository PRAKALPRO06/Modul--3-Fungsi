package main

import "fmt"

var a, b, c, d int

func faktorial(n int) int{
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) int{
	return faktorial(n) / faktorial (n-r)
}

func kombinasi(n, r int) int{
	return faktorial(n) / (faktorial (r) * faktorial (n-r))
}

func main(){
	fmt.Printf("Masukkan Input (dipisahkan oleh spasi) = ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Printf("%d, %d\n", permutasi(a, c), kombinasi(a, c))
		fmt.Printf("%d, %d\n", permutasi(b, d), kombinasi(b, d))
		
	}else {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d")
	}
}
