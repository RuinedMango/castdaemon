package main

import (
	"fmt"
	"log"
	"strings"
	"os"

	"github.com/mmcdole/gofeed"
)

func subscribe(url string){
	fp := gofeed.NewParser()
	var urlstring string = strings.Replace(url, "\n", "", -1)
	feed, err := fp.ParseURL(urlstring)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(feed.Title)
}

func writeSub(name string){
	fi, err := os.Open("~/.castd/subscribed.txt")
	if err != nil{
		log.Fatal(err)
	}
	_ = fi
}
