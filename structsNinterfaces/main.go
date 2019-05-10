package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Toggle struct {
	name    string
	off, on bool
}

// ###### Embeded types ########

type Person struct {
	Name string
}

type Android struct {
	// Person Person // this is acceptable, but not recommended
	Person // this is called an anonymous field use this way
	Model  string
}

// ########### Interfaces #################

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

var c = Circle{x: 0, y: 0, r: 5}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func newToggle(name string) Toggle {
	fmt.Println("creating a new light")
	return Toggle{
		name: name,
		off:  true,
		on:   false,
	}
}

func changeState(someSwitch *Toggle) {
	fmt.Printf("Changing the state of %s...\n", someSwitch.name)
	if someSwitch.off {
		//switch the state
		fmt.Println("Turning the switch on...")
		someSwitch.off = false
		someSwitch.on = true
	} else {
		fmt.Println("Turning the switch off...")
		someSwitch.off = true
		someSwitch.on = true
	}
	return
}

func circleArea(c *Circle) float64 {
	fmt.Printf("calculating the area of the circle...\n")
	return math.Pi * c.r * c.r
}

// method practice
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {

	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func main() {

	fmt.Println(c)
	//fmt.Println(c.y)
	//c.y = 5
	//fmt.Println(c)

	//area := circleArea(&c)
	//fmt.Println(area)
	fmt.Printf("Circle: %v\n", c)
	fmt.Println(c.area())

	light := newToggle("light_Switch")
	fmt.Println(light)

	changeState(&light)
	fmt.Println(light)

	fmt.Println("----------")
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println("-----------")
	a := new(Android)
	a.Talk()

	// interface call
	fmt.Println(totalArea(&c, &r))

	//a.talk()

}
