package main

import "fmt"

func main() {
    fmt.Println("Pointers in golang")
    //var ptr *int
    //fmt.Println("Uninitalized pointer: ",ptr)

    val := 10
    ptr := &val
    fmt.Println("Address stored in ptr: ", ptr)
    fmt.Println("Value pointed to by ptr: ", *ptr)

    *ptr *= 2

    fmt.Println("Updated value of val is: ", val)
}
