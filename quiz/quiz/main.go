package main

import (
    "os"
    "fmt" 
)

func main(){
    if len(os.Args) != 2{ 
    fmt.Println("error: Usage: ./quiz <quiz.csv>")
    fmt.Println(len(os.Args))
    return
    }

    fmt.Println("You typed -> ",os.Args[1])
    return
}
