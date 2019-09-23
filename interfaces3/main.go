package main

import (
	"fmt"
	"math"
)

//geometry type will be given to area and perimeter methods that return types of float 64
type geometry interface {
	area() float64
	perim() float64
}

//rectangle and circle templates
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//area methods, these will now also have the type of geometry
func (r rect) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//perimeter method, also now have additional geomtry type
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	//print out what is passed to this measure function, eg {3,4} rect, or {2} circle
	fmt.Println(g)
	//if the circle struct is used, then the circle area and perim methods are used, likewise if the rectangle stuct is used.
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
