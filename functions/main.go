package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Functions in golang...")
    datetime()
    fmt.Println("Addition of 2 and 3 is: ", add(2,3))
    fmt.Println("Addition of 1, 2, 3 and 4 is: ", addAll(1, 2,3,4))
}

func datetime() {
    fmt.Println("Datetime: ", time.Now())
}

func add (a int, b int) int {
    return a + b
}

// Variadic functions
func addAll(nums ...int) int {
    var sum int = 0
    // Values of variadic parameter are stored as an array
    fmt.Println("variadic parameter: ", nums) 
    for i := range nums {
        sum+=nums[i]
    }
    return sum
}
