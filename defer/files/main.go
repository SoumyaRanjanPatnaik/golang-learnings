package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("File Handling in go")
    file, err := os.Create("hello.txt")
    checkErrorNil(err)

    defer file.Close()

    
    fmt.Print("What do you want to write into the file? ")
    input, err := reader.ReadString('\n')
    checkErrorNil(err)

    _, err = io.WriteString(file, input)
    checkErrorNil(err)
    _, err = io.WriteString(file, input)
    checkErrorNil(err)

    fmt.Println("The file now contains: ", read(file))
}

func read(file *os.File) string {
    data, err := ioutil.ReadFile(file.Name())
    checkErrorNil(err)
    return string(data)
}

func checkErrorNil(err error) {
    if(err != nil) {
        panic(err)
    }
}
