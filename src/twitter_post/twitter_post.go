package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "github.com/ChimeraCoder/anaconda"
    "os"
)

type Config struct {
    ConsumerKey string `json:"consumer_key"`
    ConsumerSecret string `json:"consumer_secret"`
    AccessToken string `json:"access_token"`
    AccessTokenSecret string `json:"access_token_secret"`
}

func main(){

    // ブログのtitleを受け取る
    var title string
    fmt.Print("title: ")
    fmt.Scan(&title)

    // 投稿のフォーマットファイルを読み込み、タイトルを反映して内容の確認
    root_path := "/Users/horimasahiro/github/go/src/"
    buff_contents, err := ioutil.ReadFile(root_path+"twitter_post/post_contents.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
    contents := string(buff_contents)
    contents = fmt.Sprintf(contents, title)
    fmt.Print("\n"+contents+"\n")

    // twitterに投稿するか判断
    fmt.Println("twitterに投稿しますか? (Y/n)")
    var answer string
    fmt.Scan(&answer)
    if answer != "Y" {
        fmt.Println("投稿せずに処理を終了します。")
        os.Exit(1)
    }

    // twitterのaccess_tokenなどを設定ファイルから取得
    file, err := ioutil.ReadFile(root_path+"twitter_post/config.json")
    if err != nil {
        fmt.Println(err)
        os.Exit(-2)
    }
    var config Config
    json.Unmarshal(file, &config)

    // twitter投稿
    anaconda.SetConsumerKey(config.ConsumerKey)
    anaconda.SetConsumerSecret(config.ConsumerSecret)
    api := anaconda.NewTwitterApi(config.AccessToken, config.AccessTokenSecret)
    _, twitter_err := api.PostTweet(contents, nil)
    if twitter_err != nil {
        fmt.Println(twitter_err)
        os.Exit(-3)
    }
    fmt.Println("Twitterに投稿しました。")
}
