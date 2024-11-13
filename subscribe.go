package main

import (
	"log"
	"strings"
	"os"
	"io"
	"net/http"


	"github.com/mmcdole/gofeed"
)

func programDir() string{
	home, err := os.UserHomeDir()
	if err != nil{
		log.Fatal(err)
	}
	return home + "/.castd"
}

func subscribe(url string){
	_, err := os.Stat(programDir())
    if os.IsNotExist(err) {
        err := os.MkdirAll(programDir(), os.ModePerm)
		if err != nil{
			log.Fatal(err)
		}
    }
	_, err = os.Stat(programDir() + "/feeds")
    if os.IsNotExist(err) {
        err := os.MkdirAll(programDir() + "/feeds", os.ModePerm)
		if err != nil{
			log.Fatal(err)
		}
    }
	fp := gofeed.NewParser()
	var urlstring string = strings.Replace(url, "\n", "", -1)
	feed, err := fp.ParseURL(urlstring)
	if err != nil{
		log.Fatal(err)
	}
	writeSub(feed.Title, urlstring)
	cacheFeed(feed.Title, urlstring)
}

func writeSub(name string, feedUrl string){
	
	fi, err := os.OpenFile(programDir() + "/subscribed.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil{
		log.Fatal(err)
	}

	if _, err = fi.WriteString(name + ":" + feedUrl + "\n"); err != nil{
		log.Fatal(err)
	}
}

func cacheFeed(name string, feedUrl string){
	fi, err := os.OpenFile(programDir() + "/feeds/" + name + ".rss", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil{
		log.Fatal(err)
	}
	
	response, err := http.Get(feedUrl) //use package "net/http"

	if err != nil {
    	log.Fatal(err)
	}

	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(fi, response.Body) //use package "io" and "os"
	if err != nil {
    	log.Fatal(err)
	}

	_ = n
}
