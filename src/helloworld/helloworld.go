package main

import (
    "fmt"
    "calc"
    "github.com/bitly/go-simplejson"
)

func main(){
    fmt.Println(calc.Add(1,2))
    fmt.Println("\n")

    json := simplejson.New()
    json.Set("message", "Hello, World!")

    b, _ := json.EncodePretty()

    fmt.Printf("%s\n", b)
}
