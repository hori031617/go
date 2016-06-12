package main

import (
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string) <-chan string {
	// 関数でチャネルを生成
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}

	// チャネルを返す
	return statusChan
}

func main() {

	urls := []string{
		"http://google.com",
		"http://yahoo.co.jp",
		"http://www.nifty.com",
	}
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
