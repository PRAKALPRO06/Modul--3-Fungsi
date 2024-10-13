// Andika indra prastawa_2311102033
package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Println("Masukkan nilai a, b, c, d (pisahkan dengan spasi):")
	fmt.Scan(&a, &b, &c, &d)

	pac := permutasi(a, c)
	kac := kombinasi(a, c)

	pbd := permutasi(b, d)
	kbd := kombinasi(b, d)

	fmt.Printf("Permutasi dan kombinasi a terhadap c: P(%d, %d) = %d, C(%d, %d) = %d\n", a, c, pac, a, c, kac)
	fmt.Printf("Permutasi dan kombinasi b terhadap d: P(%d, %d) = %d, C(%d, %d) = %d\n", b, d, pbd, b, d, kbd)
}
