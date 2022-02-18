package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Switch and case in golang: ");

    rand.Seed(time.Now().UnixNano())
    diceNum := 1 + rand.Intn(6)

    fmt.Println("You got: ", diceNum)

    switch diceNum {
    case 1: 
        fmt.Println("You can open with 1")
    case 2:
        fmt.Println("Move 2 steps")
        fallthrough
    case 3: 
        fmt.Println("Move 3 steps")
        fallthrough
    case 6: 
        fmt.Println("Move 6 steps and roll again")
    default:
        fmt.Printf("Move %v steps", diceNum)
    }
}
