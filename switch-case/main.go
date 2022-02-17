package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    fmt.Println("Switch and case in golang: ");

    rand.Seed(time.Now().UnixNano())
    diceNum := rand.Intn(6)

    fmt.Println("You got: ")
}
