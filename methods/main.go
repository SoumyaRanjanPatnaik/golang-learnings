package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Method in golang: ")
    // No inheritance in golang (no parent, child, super, etc)

    soumya := User{
        "Soumya", 
        "soumya@gmail.com", 
        true, 
        20,
    }

    fmt.Println("User: ", soumya)
    fmt.Printf("User Details: %+v\n", soumya)
    fmt.Printf("User Name: %v\n", soumya.Name)
    soumya.Name = "Soumya Ranjan Patnaik"
    
    soumya.getAge()

    soumya.addAge(5)
    soumya.getAge()
    
    soumya.changeStatus()
    fmt.Printf("Modified User Details: %+v\n", soumya)

}

type User struct { // Capital U means the struct is public
    Name    string // Theese are public members are the first letter is capital
    Email   string
    Status  bool
    Age    int
}

func (u User) getAge()  {
    fmt.Printf("%v's age is %v\n", strings.TrimSpace(u.Name), u.Age)
}

// Pass by value
func (u User) addAge(x int)  {
    u.Age += x
    fmt.Printf("%v's new age is %v\n", strings.TrimSpace(u.Name), u.Age)
}

// Pass by reference
func (u *User) changeStatus () {
    u.Status = !u.Status
}
