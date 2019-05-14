package main

import "fmt"

type database map[string]float64

func (db database) check(s string) {
	if n, ok := db[s]; ok {
		fmt.Printf("%s was found!\n value: %.2f\n", s, n)
	} else {
		fmt.Printf("%s was not found, sorry..\n", s)
	}
}

func main() {
	db := database{"hat": 100}

	db["socks"] = 32

	for k, _ := range db {
		fmt.Println(k)
	}
	db.check("socks")

}
