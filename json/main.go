package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_BASE_URL = "https://jsonplaceholder.typicode.com"

type User struct {
    Name  string  `json:"name"`
    Email string  `json:"email,omitempty"`
    Id    int     `json:"userId"`
    Age   int     `json:"age"`
    Roles []string `json:"roles,omitempty"`
    Nil   []string `json:"nil"`
}

type Post struct {
    UserId    int     
    Id        int     
    Title     string  
    Body      string  
}

func main() {
    fmt.Println("Convert data to json")
    data := *new([]User)

    data = append(data, User{
        "Soumya Ranjan Patnaik",
        "soumya@google.com",
        12,
        21,
        []string{
            "Student",
            "admin",
        },
        nil,
    })
    data = append(data, User{
        "Soumya Ranjan Patnaik",
        "soumya@google.com",
        12,
        21,
        nil,
        nil,
    })

    getJson(data)
    
    strmap := []map[string]string {
        {
            "name":"Soumya",
            "age":"20",
        },
        {
            "name":"Dhruv",
            "age":"19",
        },
    }

    raw_map_json, _ := json.MarshalIndent(strmap, "", "  ")
    fmt.Println(string(raw_map_json))

    GetRequest("/posts/1")
}

func GetRequest(path string) {
    req_url := API_BASE_URL + path
    res, err := http.Get(req_url);
    if(err != nil) {
        return 
    }
    raw_bytes, err := ioutil.ReadAll(res.Body) 
    data := string(raw_bytes)
    post_obj := jsonToPost(data)
    fmt.Println("Post struct: ")
    fmt.Println("Title of post: ", post_obj.Title)
    fmt.Println("UserId: ", post_obj.UserId)
    fmt.Println("Body: ", post_obj.Body)
}

func jsonToPost (json_str string) Post {
    var post_obj Post 
    err := json.Unmarshal( []byte(json_str), &post_obj)
    if(err != nil) {
        panic(err)
    }
    return post_obj
}
func getJson(users []User ) string {
    raw_json, err := json.MarshalIndent(users, "", "  ")
    if(err != nil) {
        fmt.Errorf("Error Occurred: %v", err)
        return ""
    }
    fmt.Println("Json for struct: ", string( raw_json ))
    return string(raw_json)
}
