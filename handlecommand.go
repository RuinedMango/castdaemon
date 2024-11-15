package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func handle (buf []byte){
	command := strings.Split(string(buf), ":")[0]
	arguments := strings.SplitN(string(buf), ":", 2)[1];
	switch command{
	case "test":
		print("Testicles:")
		print(arguments)

	case "kill":
		os.Remove("/tmp/castdaemon.sock")
		os.Exit(1)

	case "subscribe":
		subscribe(arguments)

	case "download":
		//Sanitize and parse epNum
		epNum, err := strconv.ParseUint(strings.ReplaceAll(strings.Split(arguments, ":")[1], "\n", ""), 0, 32)
		if err != nil {
			log.Fatal(err)
		}
		downloadEp(strings.Split(arguments, ":")[0], epNum)

	case "delete":
		epNum, err := strconv.ParseUint(strings.ReplaceAll(strings.Split(arguments, ":")[1], "\n", ""), 0, 32)
		if err != nil{
			log.Fatal(err)
		}
		deleteEp(strings.Split(arguments, ":")[0], epNum)

	case "play":
		//Sanitze and parse epNum
		epNum, err := strconv.ParseUint(strings.ReplaceAll(strings.Split(arguments, ":")[1], "\n", ""), 0, 32)
		if err != nil{
			log.Fatal(err)
		}
		go play(strings.Split(arguments, ":")[0], epNum)

	case "stop":
		stop()

	//setup mpris compatibility
	case "mprize":
		mprize()

	//A bunch of mediaplayer commands
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
