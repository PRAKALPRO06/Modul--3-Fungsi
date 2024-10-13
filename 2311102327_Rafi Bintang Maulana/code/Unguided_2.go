package main
import (
	"fmt"
	"math"
)

func dalamLingkaranNya(cx, cy, r, x, y int) bool {
	radius := math.Sqrt(float64 ((x-cx)*(x-cx) + (y-cy)*(y-cy)) )
	return radius <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	diDalemLingkaran1 := dalamLingkaranNya(cx1, cy1, r1, x, y)
	diDalemLingkaran2 := dalamLingkaranNya(cx2, cy2, r2, x, y)

	if diDalemLingkaran1 && diDalemLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalemLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalemLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}