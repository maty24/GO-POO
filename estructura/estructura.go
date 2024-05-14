package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIds []uint
	Classes map[uint]string //le dice que el key es un uint y el value es un string

}

func (c Course) printClasses() { //esto es un método que pertenece a la estructura Course
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2]) //para quitar la coma y el espacio
}
