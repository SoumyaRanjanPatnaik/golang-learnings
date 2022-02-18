package main

import "fmt"

func main() {
    defer fmt.Println("\nDefer in golang: \n")
    defer defer1()
    defer2()

}

func defer1 () {
    defer fmt.Println("Hello")
    defer fmt.Println("World")
    fmt.Println("Soumya ")
    defer fmt.Println("Ranjan ")
    fmt.Println("Patnaik")
    defer fmt.Println("!!")
}

func defer2 () {
    for i :=0; i<5; i++ {
        defer fmt.Println(i)
    }
}
