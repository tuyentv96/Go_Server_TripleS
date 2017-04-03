package main

import (
    "fmt"
    "time"
)

func main() {
    count := 0
    timeTick := time.Tick(1 * time.Second)
    timeAfter := time.After(5 * time.Second)

    for {
        select {
        case <-timeTick:
            count++
            fmt.Printf("tick %d\n", count)

        case <-timeAfter:
            fmt.Printf("timeout\n")
            
        }
    }
}