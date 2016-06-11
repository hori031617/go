package main

import(
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Config struct{
    consumerKey string `json:"consumer_key"`
    consumerSecret string `json:"consumer_secret"`
}

func main(){

    // jsonファイルを読み込み
    root_path := "/Users/horimasahiro/github/go/src/"
    file, err := ioutil.ReadFile(root_path+"json/config.json")
    if err != nil {
        panic(err)
    }
    var config Config
    json.Unmarshal(file, &config)

    fmt.Println(config.consumerKey)
    fmt.Println(config.consumerSecret)
}
