package main

import (
	"fmt"
)

// jwtToken := 300000 // Global variables cannot use walrus operator
var jwtToken = 300000 // allowed

// First letter capital means public variable
const AccessToken string = "gibberish" // public (accessible in other files)


func main() {
    var username string = "soumya" // string
    fmt.Println(username)
    fmt.Printf("Variable is of type: %T\n",username)

    var isLoggedIn bool = false // bool 
    fmt.Println("\nLogged In: ", isLoggedIn)
    fmt.Printf("type of isLoggedIn: %T\n", isLoggedIn)

    var smallVal uint8 = 255 // 8bit unsigned int
    fmt.Println("\nSmallVal: ", smallVal)
    fmt.Printf("type: %T\n", smallVal)
    // smallVal = 256 // Error as overflow occurs

    var smallFloat float32 = 256.12312412432543 // Less precision, float64 for more
    fmt.Println("\nsmallFloat: ", smallFloat)
    fmt.Printf("type: %T\n", smallFloat)
    
    // Default values and some aliases
    var anotherVariable int 
    fmt.Println("\nanotherVariable: ", anotherVariable)
    fmt.Printf("type: %T\n", anotherVariable)


    // implicit type
    var website = "github.com"
    fmt.Println("\n", website)
    // website = 10 // Invalid as the imlicit type of the website is string

    randomNumber := 3000
    fmt.Println("\nrandom Number: ", randomNumber)
    fmt.Println("\nAccessToken: ",AccessToken) //Accessing token

}
