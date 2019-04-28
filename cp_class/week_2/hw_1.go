package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var n int

	//don't use range here, you don't want the first arg!

	for i := 1; i < len(os.Args); i++ {
		fn := os.Args[i]
		text, err := ioutil.ReadFile(fn)

		// handle the case of a bad file name

		if err != nil {
			fmt.Fprintf(os.Stderr, "can't read %s: %s\n", fn, err)

			continue
		}
		// magic haooens here
		// we must convert the []byte from ReadFile]

		words := strings.Fields(string(text))
		n += len(words)
	}

	fmt.Println(n, "total words")

}
