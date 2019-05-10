package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Toggle struct {
	name    string
	off, on bool
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

var c = Circle{x: 0, y: 0, r: 5}

func circleArea(c *Circle) float64 {
	fmt.Printf("calculating the area of the circle...\n")
	return math.Pi * c.r * c.r
}

func main() {

	fmt.Println(c)
	fmt.Println(c.y)
	c.y = 5
	fmt.Println(c)

	area := circleArea(&c)
	fmt.Println(area)

	light := newToggle("light_Switch")
	fmt.Println(light)

	changeState(&light)
	fmt.Println(light)

}
