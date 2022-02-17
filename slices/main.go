package main

import (
	"fmt"
	"sort"
)

func main() {
    fmt.Println("Slices in go")
    //fruitList := []string{} // empty slice
    fruitList := []string{"Apple", "Tomato", "Peach"} // empty slice

    fmt.Printf("Type of fruitList is %T\n", fruitList)
    fmt.Println("fruitList: ", fruitList)
    fmt.Println("Size of fruitList: ", len(fruitList))

    fruitList = append(fruitList, "avacardo", "pineapple")
    fmt.Println("Updated fruitList: ", fruitList)
    fmt.Println("Size of fruitList: ", len(fruitList))

    // Slicing a slice
    fmt.Println("Slicing fruitList 1 3", fruitList[1:3]) // Last index not inclusive

    fruitList = fruitList[:4]
    fmt.Println("Updated fruitList: ", fruitList)
    fmt.Println("Size of fruitList: ", len(fruitList))

    // Allocate slice using make keyword
    scores := make([]int, 5)  // uninitialized
    fmt.Println("scores: ", scores)

    scores[0] = 8
    scores[1] = -6
    scores[2] = 6
    scores[3] = 2
    scores[4] = 11
    // scores[5] = 123 // Panic as index out of range
    fmt.Println("initialized scores: ", scores)
    scores = append(scores, 24,69, 32)
    fmt.Println("Updated Scores: ", scores)
    fmt.Println("Is Sorted? ", sort.IntsAreSorted(scores))
    sort.Ints(scores)
    fmt.Println("sorted Scores: ", scores)
    fmt.Println("Is Sorted? ", sort.IntsAreSorted(scores))
}
