package main

import "fmt"

func main() {
    var a, b, c, d int
    fmt.Print("Masukkan nilai (a b c d): ")
    fmt.Scan(&a, &b, &c, &d)

    if a >= c && b >= d {
        fmt.Println("Permutasi (a, c):", permutasi(a, c))
        fmt.Println("Kombinasi (a, c):", kombinasi(a, c))
        fmt.Println("Permutasi (b, d):", permutasi(b, d))
        fmt.Println("Kombinasi (b, d):", kombinasi(b, d))
    } else {
        fmt.Println("Tidak memenuhi kondisi")
    }
}

func faktorial(n int) int {
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
