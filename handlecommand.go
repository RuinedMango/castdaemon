package main

import "os"

func handle (buf []byte){
	switch buf[0]{
	case 'a':
		print("first A-ide")
	case 'x':
		os.Remove("/tmp/castdaemon.sock")
		os.Exit(1)
	}
}
