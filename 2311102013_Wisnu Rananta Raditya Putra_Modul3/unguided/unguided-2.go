package main

import "fmt"

func f(x_2311102013 int) int {
    return x_2311102013 * x_2311102013
}

func g(x_2311102013 int) int {
    return x_2311102013 - 2
}

func h(x_2311102013 int) int {
    return x_2311102013 + 1
}

func fogoh(x_2311102013 int) int {
    return f(g(h(x_2311102013)))
}

func gohof(x_2311102013 int) int {
    return g(h(f(x_2311102013)))
}

func hofog(x_2311102013 int) int {
    return h(f(g(x_2311102013)))
}

func main() {
    var a, b, c int

    fmt.Print("Masukkan Tiga Bilangan Bulat a, b ,c : ")
    fmt.Scan(&a, &b, &c)

    fmt.Println(fogoh(a))
    fmt.Println(gohof(b))
    fmt.Println(hofog(c))
}