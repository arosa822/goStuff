package main

import "fmt"

type Point struct {
	name string
	x, y int
}

type Shape interface {
	area() float64
}

// initialize the structure
//var c Point

// another approach is to use shorthand ->  `c := new(Point)`

func (c *Point) update(x, y int) {
	c.x = x
	c.y = y
}

func (c *Point) displayVal() {
	fmt.Printf("name:%s, x:%d, y:%d\n", c.name, c.x, c.y)
}

func main() {
	c := new(Point)
	c.displayVal()
	a := Point{name: "a", x: 1, y: 2}
	a.displayVal()
	a.update(10, 11)
	a.displayVal()
}
