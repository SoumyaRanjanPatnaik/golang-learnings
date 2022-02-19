package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_BASE_URL = "https://jsonplaceholder.typicode.com"

func main() {
    fmt.Println("GET request...")
    res, _ := GetRequest("/posts/1")
    fmt.Println("Response: ", res)
}

func GetRequest(path string) (data string, err error) {
    req_url := API_BASE_URL + path
    res, err := http.Get(req_url);
    if(err != nil) {
        return "" , err
    }
    raw_bytes, err := ioutil.ReadAll(res.Body) 
    data = string(raw_bytes)
    return data, err
}
