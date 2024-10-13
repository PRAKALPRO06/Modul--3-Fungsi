package main

import (
	"fmt"
)

func fungsiF (x int) int {
	return (x * x)
}
func fungsiG (x int) int {
	return (x - 2)
}
func fungsiH (x int) int {
	return (x + 1)
}

func komposisiA (x int) int {
	return fungsiF(fungsiG(fungsiH(x)))
}
func komposisiB (x int) int {
	return fungsiG(fungsiH(fungsiF(x)))
}
func komposisiC (x int) int {
	return fungsiH(fungsiF(fungsiG(x)))
}

func main()  {
	var a, b, c int

	fmt.Print("Masukkan input = ")
	fmt.Scan(&a, &b, &c)

	A := komposisiA(a)
	B := komposisiB(b)
	C := komposisiC(c)

	fmt.Print("\n", A)
	fmt.Print("\n", B)
	fmt.Print("\n", C)
}