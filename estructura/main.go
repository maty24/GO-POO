package main

import "fmt"

func main() {
	Go := Course{
		Name:    "Go desde cero",
		Price:   12.34,
		IsFree:  false,
		UserIds: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	js := Course{}
	js.Name = "JS desde cero"
	js.Price = 34.56
	js.UserIds = []uint{12, 56}

	fmt.Println(Go.Name)
	fmt.Println(css.Name)
	fmt.Println(js.Name)

	Go.printClasses()
}
