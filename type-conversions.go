package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z float64 = f
	fmt.Println(x, y, f, z)
}
