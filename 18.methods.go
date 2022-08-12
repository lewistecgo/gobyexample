package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 20}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimetro: ", r.perim())

	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimetro: ", rp.perim())

}