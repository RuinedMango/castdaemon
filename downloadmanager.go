package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mmcdole/gofeed"
)

func downloadDir() string{
	return programDir() + "/downloads"
}

func downloadEp(podcastName string, epNum uint64){
	_, err := os.Stat(downloadDir() + "/" + podcastName)
    if os.IsNotExist(err) {
        err := os.MkdirAll(downloadDir() + "/" + podcastName, os.ModePerm)
		if err != nil{
			log.Fatal(err)
		}
    }
	fp := gofeed.NewParser()
	content, err := os.ReadFile(programDir() + "/feeds/" + podcastName + ".rss")
    if err != nil {
        log.Fatal(err)
    }
	feed, err := fp.ParseString(string(content))
	ep := feed.Items[uint64(len(feed.Items) - 1) - (epNum - 1)]
	epurl := ep.Enclosures[0].URL
	err = DownloadFile(downloadDir() + "/" + podcastName + "/" + strconv.FormatUint(epNum, 32) + ".mp3", epurl)
	if err != nil{
		log.Fatal(err)
	}
}

func deleteEp(podcastName string, epNum uint64){
	err := os.Remove(downloadDir() + "/" + podcastName + "/" + strconv.FormatUint(epNum, 32) + ".mp3")
	if err != nil{
		log.Fatal(err)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
