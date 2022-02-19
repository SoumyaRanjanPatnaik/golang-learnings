package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const BASE_URL = "https://jsonplaceholder.typicode.com/"

func main() {
    fmt.Println("Post request: ")
    PostRequest("posts", `{
        "title": "Hello World",
        "body": "I'm Soumya Ranjan Patnaik, learning golang",
        "userId": 1
    }`)
}

func PostRequest(path string, data string) {
    req_url := BASE_URL + path
    fmt.Println("POST 1.1 /", req_url)
    payload := strings.NewReader(data)
    const content_type = "applicationo/json"
    res, err := http.Post(req_url, content_type, payload);
    if(err != nil) {
        return
    }
    defer res.Body.Close()
    fmt.Println("Successfully Created Resource:")
    res_bytes, err := ioutil.ReadAll(res.Body)
    if(err != nil) {
        return
    }
    fmt.Println(string(res_bytes))
}
