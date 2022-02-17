package main

import "fmt"

func main() {
    fmt.Println("Structs in golang: ")
    // No inheritance in golang (no parent, child, super, etc)

    soumya := User{"Soumya", "soumya@gmail.com", true, 20}
    fmt.Println("User: ", soumya)
    fmt.Printf("User Details: %+v\n", soumya)
    fmt.Printf("User Name: %v", soumya.Name)
}

type User struct { // Capital U means the struct is public
    Name    string // Theese are public members are the first letter is capital
    Email   string
    Status  bool
    Age    int
}
