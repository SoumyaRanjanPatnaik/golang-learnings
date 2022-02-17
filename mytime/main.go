package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Time in golang: \n")
    presentTime := time.Now();
    fmt.Println(presentTime)
    fmt.Println(presentTime.Format("01-02-2006 Mon"))
    newDate  := time.Date(2022, time.February, 17, 04, 23, 11, 1,time.Local)
    fmt.Println("New Date: ", newDate)
}
