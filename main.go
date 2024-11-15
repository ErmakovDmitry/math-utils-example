package main

import (
	"fmt"
	"strconv"

	"github.com/ErmakovDmitry/math-utils/conv"
)

func main() {
	var r1 float64 = conv.Add(1.4, 2.3, 8.7)

	fmt.Println("Float64ToString(r1) = " + Float64ToString(r1))
	fmt.Printf("r1 = %f\n", r1)

	r2 := fmt.Sprintf("%f", r1)
	fmt.Printf("r2 = %s\n", r2)
}

func Float64ToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}

func Float32ToString(input_num float32) string {
	return strconv.FormatFloat(float64(input_num), 'f', -1, 32)
}
