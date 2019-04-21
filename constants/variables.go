package main

import "fmt"

const (
    Pi = 3.14
    Truth = false
    Big = 1 << 62
    Small = Big >> 61
)

var pi float32  = 3223.14

func main(){
    const Greetings = "hello"
    //fmt.Println(Pi)
    //fmt.Println(Truth)
    //fmt.Println(Big)
    //fmt.Println(Small)
    //fmt.Println(Greetings+ "\n")
    fmt.Printf("%s \n big = %d \n small = %d\n", Greetings, Big, Small)

    cycloneModel := 6
    fmt.Println(cycloneModel)
    name := "caprica-six"
    aka := fmt.Sprintf("Number %d",6)
    fmt.Printf("%s is also a known as %s\n",
        name, aka)
    //pi = pi + 2
    pi = Pi + 1
    fmt.Println(pi)
}
