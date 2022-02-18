package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const base_url = "https://soumyaranjanpatnaik.me"

func main() {
    fmt.Println("Web request in web request")

    resp, err := http.Get(base_url)
    checkErrorNil(err)
    defer resp.Body.Close();

    fmt.Printf("Response is of type: %T\n", resp)

    data, err := ioutil.ReadAll(resp.Body)
    checkErrorNil(err)
    
    fmt.Println(string(data))
}

func checkErrorNil (err error) {
    if(err != nil) {
        panic(err)
    }
}
