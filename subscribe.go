package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mmcdole/gofeed"
)

func subscribe(url []byte){
	fp := gofeed.NewParser()
	var urlstring string = strings.Replace(string(url[:]), "\n", "", -1)
	feed, err := fp.ParseURL(urlstring)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(feed.Title)
}
