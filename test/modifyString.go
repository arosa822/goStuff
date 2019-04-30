package main

import (
    "fmt"
    "os"
    "strings" 
)

func checkAndModify(s string) string{
    
    //l := len(os.Args)

    if strings.HasPrefix(s,"www") == true {
        fmt.Println("YES!\n")
    } else {
        s = "www."+ s
        fmt.Println("NO!\n")
    }
    return s
}
func main() {
    input := os.Args[1]
    fmt.Printf(" %s\n",checkAndModify(input))
}
