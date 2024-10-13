// 2311102018- Aryo Tegar Sukarno
package main

import (
        "fmt"
        "math/big"
)

func factorial(n int) *big.Int {
        if n == 0 {
                return big.NewInt(1)
        }
        return big.NewInt(int64(n)).Mul(big.NewInt(int64(n)), factorial(n-1))
}

func permutasi(n, r int) *big.Int {
        return factorial(n).Div(factorial(n), factorial(n-r))
}

func kombinasi(n, r int) *big.Int {
        return factorial(n).Div(factorial(n), factorial(r).Mul(factorial(r), factorial(n-r)))
}

func main() {
        var a, b, c, d int
        fmt.Scan(&a, &b, &c, &d)

        // Periksa syarat a >= c dan b >= d
        if a < c || b < d {
                fmt.Println("Input tidak valid: a harus lebih besar sama dengan c, dan b harus lebih besar sama dengan d")
                return
        }

        // Hitung permutasi dan kombinasi untuk a dan c
        p1 := permutasi(a, c)
        c1 := kombinasi(a, c)

        // Hitung permutasi dan kombinasi untuk b dan d
        p2 := permutasi(b, d)
        c2 := kombinasi(b, d)

        // Tampilkan hasil
        fmt.Println(p1, c1)
        fmt.Println(p2, c2)
}
