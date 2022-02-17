// Package main provides ...
package main

import "fmt"

func main() {
    fmt.Println("Maps in golang:\n")
    mymap := make(map[string] string)
    mymap["Hello"] = "world"
    mymap["js"] = "javascript"
    fmt.Println("mymap: ", mymap)
    fmt.Println("Hello: ", mymap["Hello"])

    // Delete value
    delete(mymap, "Hello")
    fmt.Println("map after deleting Hello from mymap", mymap)
    
    //Random data
    mymap["soumya"] = "twenty"
    mymap["tesla"] = "car"

    fmt.Println("\nKey value pairs for my map:")
    for key, val := range mymap {
        fmt.Printf("Key: %v   \tVal: %v\n", key, val)
    }
}
