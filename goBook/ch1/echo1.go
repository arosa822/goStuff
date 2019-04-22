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

// create methods to print out each command line argument 
func main() {
    // Method 1 - explicit
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	fmt.Println(countArgs())
	//fmt.Println(s)

    // Method 2 implicit 
    // s, sep := "", ""
    // for _, arg := range os.Args[1:] {
    //     s += sep + arg
    //     sep = " "
    // }
    // fmt.Println(s)

    // method 3 print values
    fmt.Println(strings.Join(os.Args[1:]," "))
    fmt.Println(os.Args[1:])

    // Exersizes: 
    // 1.1  Modify to print out os.args[0] the command
    // 1.2 print index and value of each argument per line
    // 1.3 benchmark the different methods above see 1.6 (time) and 11.4
    // benchmark tests
}
