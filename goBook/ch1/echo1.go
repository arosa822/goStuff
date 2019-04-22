// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

// this function displays the number of args
func countArgs() string {
	var num int = len(os.Args)
	message := fmt.Sprintf("%d arguments passed through!\n", num)

	return message
}
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(countArgs())
	fmt.Println(s)
}
