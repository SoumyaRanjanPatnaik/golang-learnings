package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const API_BASE_URL = "jsonplaceholder.typicode.com"

func main() {
    fmt.Println("GET request...")
    res, _ := GetRequest("/posts/1")
    fmt.Println("Response: ", res)
}

func GetRequest(path string) (data string, err error) {
    defer func () {
        if err := recover(); err != nil {
            panic(err)
        }
    } ()
    request_url := &url.URL{
        Scheme: "https",
        Host: API_BASE_URL,
        Path: path,
        
    }
    fmt.Println(request_url.String())
    resp, err := http.Get(request_url.String())
    if err != nil {
        return "" , err
    }
    data_bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "" , err
    }
    data = string(data_bytes)
    return data, err
}
