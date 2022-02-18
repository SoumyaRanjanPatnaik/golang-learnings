package main

import (
	"fmt"
	"net/url"
)

const my_url = "http://soumyaranjanpatnaik.me:3000/resume?section=skills&curr=golang"
func main() {
    fmt.Println("Urls in golang")
    fmt.Println("URL: ")

    // Parsing url
    parsedUrl, _ := url.Parse(my_url)
    fmt.Println(parsedUrl.Scheme)
    fmt.Println(parsedUrl.Host)
    fmt.Println(parsedUrl.Path)
    fmt.Println(parsedUrl.Port())
    fmt.Println(parsedUrl.RawQuery)

    qparams := parsedUrl.Query()
    fmt.Println("params: ",qparams)

    // values obtained are string arrays
    fmt.Printf("Type of value in qparams map is: %T\n", qparams["curr"])

    for key, val := range qparams {
        fmt.Println(key,": ", val[0])
    }
    
    partsOfUrl := &url.URL {
        Scheme: "https",
        Host: "soumyaranjanpatnaik.me",
        Path: "/blog",
        RawQuery: "q=auth",
    }

    fmt.Println("Constructed url: ", partsOfUrl.String())

}
