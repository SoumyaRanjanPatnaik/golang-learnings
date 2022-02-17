package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the rating: ")
    input, _ := reader.ReadString('\n')
    rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
        fmt.Println(err)
    } else{
        fmt.Println("Rating +1 : ", rating + 1)
    }
}
