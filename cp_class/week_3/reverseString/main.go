package main

import "fmt"

func reverse(s []int) []int {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
	return s
}

func revPoint(s *[4]int) {
	for a, b := 0, len(*s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
}

func main() {
	sl := [...]int{0, 1, 2, 3, 4}
	sj := [...]int{5, 6, 7, 8}

	fmt.Println(reverse(sl[:]))
	revPoint(&sj)
	fmt.Println(sj)

}
