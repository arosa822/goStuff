package main

import "fmt"

func makeLine() {
	fmt.Println("-------------------------------------")
}

//nonempty returns a slice holding only the non-empty strings.
// the underlying array is modified during the call..
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // Zero Length slice of the original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reversePointer(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Println("nonemty")
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	makeLine()

	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2))
	fmt.Printf("%q\n", data2)
	makeLine()

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

}
