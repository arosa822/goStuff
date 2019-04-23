// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

// this function displays the number of args
func countArgs() string {
	var num int = len(os.Args)
	message := fmt.Sprintf("%d arguments passed through!\n", num)

	return message
}

func echo1() string {
	// initialize 2 vars with default values ""
	var s, sep string
	// start loop (method 1)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	// initialize 2 vars and explicitly assign values
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	fmt.Println(countArgs())
	//fmt.Println(echo1())
	//fmt.Println(echo2())
	//fmt.Println(echo3())
	fmt.Println(os.Args[0:])

	// Exersizes:
	// 1.1 Modify to print out os.args[0] the command
	// 1.2 print index and value of each argument per line
	// 1.3 benchmark the different methods

	//1.2

	// start the loop
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}

}
