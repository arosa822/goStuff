package main

import "fmt"

const (
    Pi = 3.14
    Truth = false
    Big = 1 << 62
    Small = Big >> 61
)

func main(){
    const Greetings = "hello" 
    fmt.Println(Pi)
    fmt.Println(Truth)
    fmt.Println(Big)
    fmt.Println(Small)
    fmt.Println(Greetings)
}
