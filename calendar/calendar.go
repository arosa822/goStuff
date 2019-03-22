package main

import "fmt"
import "time"

func main() {
    t1 := Date(2019, 07, 13)
    t2 := Date(2019, 10, 14)
    days := t2.Sub(t1).Hours() / 24
    fmt.Println(days)
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
