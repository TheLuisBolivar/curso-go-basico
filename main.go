package main

import (
	"fmt"
	"math/cmplx"
)

const (
	ConstX = 43
	ConstY = 45.5
	ConstU = "hello world"
)

var (
	ToBe       bool       = false
	MaxInt     uint64     = 1<<61 - 1
	z          complex128 = cmplx.Sqrt(-5 + 12i)
	w          int
	y          bool
	s          string
	pruebaCast float64 = 64.43
)

//go run main.go
func main() {
	i := int(pruebaCast)
	fmt.Printf("Type: %T value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v \n", z, z)
	fmt.Printf("Type: %T value: %v \n", w, w)
	fmt.Printf("Type: %T value: %v \n", y, y)
	fmt.Printf("Type: %T value: %v \n", s, s)
	fmt.Printf("Type: %T value: %v \n", pruebaCast, pruebaCast)
	fmt.Printf("Type: %T value: %v \n", i, i)
	fmt.Printf("Type: %T value: %v \n", ConstX, ConstX)
	fmt.Printf("Type: %T value: %v \n", ConstY, ConstY)
	fmt.Printf("Type: %T value: %v \n", ConstU, ConstU)
}
