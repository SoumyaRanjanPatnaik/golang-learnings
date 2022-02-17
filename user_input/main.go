package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    welcome := "WELCOME"
    fmt.Println(welcome)
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the rating: ")
    // Comma ok syntax (first variable stores result and the last stores error)
    rating, _ := reader.ReadString('\n') // _ means the error is not used
    fmt.Println("You rated us:", rating)
    fmt.Printf("Typeof rating: %T\n", rating) 
}
