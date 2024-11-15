package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/mmcdole/gofeed"
)

var globalformat beep.Format
var globalstreamer beep.StreamSeekCloser

func play(podcastName string, epNum uint64) {
	f, err := os.Open(downloadDir() + "/" + podcastName + "/" + strconv.FormatUint(epNum, 32) + ".mp3")
	var streamer beep.StreamSeekCloser
	var format beep.Format
	if os.IsNotExist(err){
		streamer, format, err = mp3.Decode(playweb(podcastName, epNum))
		if err != nil{
			log.Fatal(err)
		}
	} else {
		streamer, format, err = mp3.Decode(f)
		if err != nil{
			log.Fatal(err)
		}
	}
	globalstreamer = streamer
	globalformat = format
	defer streamer.Close()

	sr := format.SampleRate * 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	done := make(chan bool)
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func playweb(podcastName string, epNum uint64) io.ReadCloser{
	fp := gofeed.NewParser()
	content, err := os.ReadFile(programDir() + "/feeds/" + podcastName + ".rss")
    if err != nil {
        log.Fatal(err)
	}
	feed, err := fp.ParseString(string(content))
	ep := feed.Items[uint64(len(feed.Items) - 1) - (epNum - 1)]
	epurl := ep.Enclosures[0].URL
	resp, err := http.Get(epurl)
	if err != nil{
		log.Fatal(err)
	}
	returner := resp.Body
	return returner
}

func stop(){
	speaker.Close()
}
