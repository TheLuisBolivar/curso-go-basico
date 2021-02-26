package main

import "fmt"

//go run main.go
func main() {
	i, j, x := 5, "hello", true

	p := &i
	fmt.Printf("p is: %T and his value is: %v \n", *p, *p)
	*p = 100
	fmt.Printf("cambie el valor de i por medio de derefenciacion del puntero p: %T %v \n", i, i)

	m := &j
	fmt.Printf("m is: %T and his value is: %v \n", *m, *m)
	*m = "world"
	fmt.Printf("cambie el valor de j por medio de derefenciacion del puntero m: %T %v \n", j, j)

	n := &x
	fmt.Printf("n is: %T and his value is: %v \n", *n, *n)
	*n = false
	fmt.Printf("cambie el valor de x por medio de derefenciacion del puntero n: %T %v \n", x, x)
}
