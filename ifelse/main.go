package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("If Else in golang")
    marks :=23
    
    if marks <= 33 {
        fmt.Println("Failed with marks: ", marks)
    } else if marks >90 {
        fmt.Println("Passed with good marks")
    } else {
        fmt.Println("Passed")
    }

    if 5%2 == 0 {
        fmt.Println("5 is even")
    } else {
        fmt.Println("5 is odd")
    }

    if num := 10; num == 10 {
        fmt.Println("Perfect cgpa")
    } else {
        fmt.Println("CGPA is: ", num)
    }

    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if(err != nil) {
        fmt.Println("Error occurred")
    } else {
        fmt.Println("Success, input: ", input)
    }

}
