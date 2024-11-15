package main

import (
	"fmt"

	"github.com/ErmakovDmitry/math-utils/calc"
	"github.com/ErmakovDmitry/math-utils/conv"
)

func main() {
	var r1 float64 = calc.Add64(1.4, 2.3, 8.7)

	fmt.Println("Float64ToString(r1) = " + conv.Float64ToString(r1))
	fmt.Printf("r1 = %f\n", r1)

	r2 := fmt.Sprintf("%f", r1)
	fmt.Printf("r2 = %s\n", r2)

	var s1 = calc.Sub64(3.7)
	fmt.Printf("s1 = %f\n", s1)

	var r32 float32 = calc.Add32(1.4, 2.3, 8.7)
	fmt.Printf("r32 = %f\n", r32)
	fmt.Println("Float32ToString(r32) = " + conv.Float32ToString(calc.Sub32(r32, 0.4)))
}
