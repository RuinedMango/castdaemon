package main

import (
	"time"

	"github.com/gopxl/beep/v2/speaker"
)

var paused = false
func pause(){
	speaker.Lock()
	paused = true
}

func unpause(){
	speaker.Unlock()
	paused = false
}

func pausetoggle(){
	if paused{
		speaker.Unlock()
		paused = false
	}else{
		speaker.Lock()
		paused = true
	}
}

func forwardskip(){
	print("Forwardskip Requested")
}

func backwardskip(){
	print("Backwardskip Requested")
}

func forward(){
	globalstreamer.Seek(globalformat.SampleRate.N(globalformat.SampleRate.D(globalstreamer.Position()).Round(time.Second) + (time.Second * 5)))
	print("Forward Requested")	
}

func backward(){
	globalstreamer.Seek(globalformat.SampleRate.N(globalformat.SampleRate.D(globalstreamer.Position()).Round(time.Second) - (time.Second * 5)))
	print("Backward Requested")
}
