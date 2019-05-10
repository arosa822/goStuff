package main

import "fmt"

var z int = 3

func add(a int) int {

	a = a + 1
	return a
}

func main() {
	i, j := 11, 2335

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println("------------")

	n := &z
	m := add(*n)
	fmt.Println(m)
}
