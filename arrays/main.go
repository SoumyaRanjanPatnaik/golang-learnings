package main

import "fmt"

func main() {
    fmt.Println("Arrays in golang: ")
    
    var fruitList [4]string
    fmt.Println("Empty array: ", fruitList) // Space printed for every element

    fmt.Printf("Type of fruitList array is: %T\n", fruitList)
    
    fruitList[0] = "Apple"
    fruitList[1] = "Orange"
    fruitList[3] = "Banana"
    fmt.Println("FruitList: ", fruitList) // space printed for second element (non initialized)

    var numList [4]int

    fmt.Println("numList: ", numList) // uinitialized value is 0
    numList[0] = 19
    numList[1] = 21
    numList[3] = 23
    fmt.Println("numList: ", numList)

    fmt.Println("Length of numList: ", len(numList))

    // Initializer for array
    anotherArray := [10]int{1,2,3} // define and initialized at the saame time
    fmt.Println("another array:", anotherArray)

    slice := anotherArray[1:4]
    fmt.Println("Slice to array: ", slice)

    slice = append(slice, 2)
    fmt.Println("anotherArray: ", anotherArray)
}
