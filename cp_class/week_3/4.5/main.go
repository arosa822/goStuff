package main

import "fmt"

func removedup(strings []string) []string {
	last := ""
	offset := 0
	for n, s := range strings {
		if s == "" {
			continue
		}
		if last == s {
			offset++
		}

		strings[n-offset] = s
		last = s
	}
	return strings[:len(strings)-offset]
}

func main() {

	a := []string{"one", "one", "two", "three"}
	fmt.Println("\nCharacteristics of a:")
	fmt.Printf("Len: %d\nType: %T\n\n", len(a), a)
	fmt.Println(removedup(a))

}
