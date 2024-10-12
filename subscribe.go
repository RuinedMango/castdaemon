package main

import(
	"github.com/mmcdole/gofeed"
	"fmt"
)

func subscribe(url []byte){
	fp := gofeed.NewParser()
	print(string(url[:]))
	feed, _ := fp.ParseURL("https://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)
}
