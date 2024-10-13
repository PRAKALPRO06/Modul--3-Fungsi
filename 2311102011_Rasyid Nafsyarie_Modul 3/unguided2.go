package main

import "fmt"

func main() {
        var a_2311102011, b, c int

        fmt.Scan(&a_2311102011, &b, &c)

        fmt.Println(fogoh(a_2311102011))
        fmt.Println(gohof(b))
        fmt.Println(hofog(c))
		
}

func f(x int) int {
        return x * x
}

func g(x int) int {
        return x - 2
}

func h(x int) int {
        return x + 1
}

func fogoh(x int) int {
        return f(g(h(x)))
}

func gohof(x int) int {
        return g(h(f(x)))
}

func hofog(x int) int {
        return h(f(g(x)))
}