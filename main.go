package main

import "fmt"

func main ()  {
	var x int = 100
	y := 10

	//Esto es un comentario

	/*
		Todo lo que esta dentro de esto,
		Se vuelven comentarios, 
		hasta parrafos, si en algun momento son necesarios 
	*/

	var boolX bool = true
	boolY := false

	fmt.Println(boolX || boolY)
	fmt.Println(boolX && boolY)
	fmt.Println(!boolX)
	fmt.Println(!boolY)

	fmt.Println(x / y)
	funcionPrivada()
	FuncionPublica()
	imprimirHola()
}

func funcionPrivada(){
	fmt.Println("Esto es una función privada perritos.")
}

func FuncionPublica(){
	fmt.Println("Esto es una función pública perritos.")
}

func imprimirHola(){
	defer fmt.Println("Estoy al comienzo, pero me imprimo al final.")
	fmt.Println("Estoy al final.")
}