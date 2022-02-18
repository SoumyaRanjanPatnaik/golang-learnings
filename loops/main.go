package main

import "fmt"

func main() {
    fmt.Println("Loops in golang: ")
    days := []string{"Sunday", "Monday",
    "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    
    // Classic for loop
    for i:=0; i<len(days); i++ {
        fmt.Println(days[i])
    }

    fmt.Println("")
    // For with range
    for i:= range days {
        fmt.Println(days[i])
    }

    fmt.Println("")
    
    // For with index and value
    for index, value := range days {
        fmt.Printf("index: %v, value: %v\n", index, value)
    }

    fmt.Println("")
    // While loop style for loop
    count := 0
    for count < 10 {
        //if(count ==8){
            //break
        //}
        if(count ==8){
            goto exit
        }
        if (count == 5) {
            count++
            continue
        }
        fmt.Println(count)
        count++
    }
    exit:
        fmt.Println("Jumped using goto")
}
