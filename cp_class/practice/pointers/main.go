package main

import "fmt"

var a int = 2

func add(b int) {
	a = a + b
}

func main() {
	b := 3
	var c int
	add(b)
	fmt.Printf("a: %d\nb: %d\nc: %d\n", a, b, c)
}
