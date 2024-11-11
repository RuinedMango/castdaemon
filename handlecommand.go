package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func handle (buf []byte){
	command := strings.Split(string(buf), ":")[0]
	arguments := strings.Split(string(buf), ":")[1];
	switch command{
	case "test":
		print("Testicles")
	case "kill":
		os.Remove("/tmp/castdaemon.sock")
		os.Exit(1)
	case "subscribe":
		subscribe(arguments)
	case "download":
		epNum, err := strconv.ParseUint(strings.Split(arguments, ":")[1], 0, 64);
		if err != nil {
			log.Fatal(err)
		}
		download(strings.Split(arguments, ":")[0], epNum)
	case "play":
		epNum, err := strconv.ParseUint(strings.Split(arguments, ":")[1], 0, 64)
		if err != nil{
			log.Fatal(err)
		}
		play(strings.Split(arguments, ":")[0], epNum)
	case "stop":
		stop()
	case "mprize":
		mprize()
	case "pause":
		pause()
	case "unpause":
		unpause()
	case "ptoggle":
		pausetoggle()
	case "forward":
		forward()
	case "backward":
		backward()
	case "fskip":
		forwardskip()
	case "bskip":
		backwardskip()
	}
}
