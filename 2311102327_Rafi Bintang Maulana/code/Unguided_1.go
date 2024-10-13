package main
import (
	"fmt"
)

func FungsiFx(x int) int {
	return x * x
}

func FungsiGx(x int) int {
	return x - 2
}

func FungsiHx(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return FungsiFx(FungsiGx(FungsiHx(x)))
}

func gohof(x int) int {
	return FungsiGx(FungsiHx(FungsiFx(x)))
}

func hofog(x int) int {
	return FungsiHx(FungsiFx(FungsiGx(x)))
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}