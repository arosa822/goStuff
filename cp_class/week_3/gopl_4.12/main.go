package main

import (
	"crypto/sha256"
	"fmt"
)

var a [3]int // array of 3 integers
var q [3]int = [3]int{1, 2, 3}

func sha256Exe() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func printSlices() {
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]
	for i, v := range q {
		fmt.Printf("%d,%d\n", i, v)
	}
}

func main() {
	printSlices()
	sha256Exe()
}
