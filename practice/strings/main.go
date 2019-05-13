package main

import "fmt"

var s1 string = "DR" // (-1,+1)

var s2 string = "UL"

type Move struct {
	x, y int
}

func parse(s string) []string {
	var parsed []string
	for _, i := range s {
		//fmt.Printf("%d:%s\n", n, string(i))
		parsed = append(parsed, string(i))
	}
	return parsed
}

func (m *Move) moveAccordingly(input []string) {

	for _, s := range input {
		if s == "U" { // U
			m.y--
		}
		if s == "D" { // D
			m.y++
		}
		if s == "L" { // L
			m.x--
		}
		if s == "R" { // R
			m.x++
		}
	}
}

func main() {
	loc := Move{x: 0, y: 0}
	fmt.Println(loc.x, loc.y)

	move := parse(s1)
	fmt.Println(move)

	loc.moveAccordingly(parse(s1)) //(-1 +1)
	fmt.Println(loc)

	move = parse(s2)
	fmt.Println(move)

	loc.moveAccordingly(parse(s2)) // (
	fmt.Println(loc)

	//for _, s := range s1 {
	//	fmt.Println(s)
	//}

}
