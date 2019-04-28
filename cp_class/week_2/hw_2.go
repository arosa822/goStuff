package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	m := make(map[string]int) // can't just use var

	if len(os.Args) > 1 {

		for i := 1; i < len(os.Args); i++ {
			fn := os.Args[i]
			text, err := ioutil.ReadFile(fn)

			if err != nil {
				fmt.Fprintf(os.Stderr, "ReadFile: %v\n", err)
			}

			words := strings.Fields(string(text))

			for _, w := range words { //ignore keys
				m[w] += 1
			}
		}
		fmt.Println(len(m), "unique words")

	} else {
		fmt.Println("no file was specified to parse")
	}

}
