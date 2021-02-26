package main

import "fmt"

type Vertex struct {
	I    int
	Y    int
	Name string
}

type Person struct {
	Name, Lastname string
	Id             int
}

var (
	mono    = Person{"Luis", "Bolivar", 1023}
	juancho = Person{"Juan", "Bolivar", 12312}
)

func main() {
	vertex := Vertex{1, 2, "name"}
	fmt.Printf("Esta es mi primer estructura: Vertex structs: %T, %v", vertex, vertex)

	p := &vertex
	p.Name = "Nuevo nombre!"
	fmt.Printf("Revisando mi puntero p :D --> %T %v", p, p)

	e := vertex.Name
	fmt.Printf("estoy revisando e %T %v", e, e)

	fmt.Println(mono.Id, juancho)
}
