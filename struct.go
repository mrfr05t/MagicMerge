package main

import "flag"

var (
	useNgrok  = flag.Bool("ngrok", false, "Set to true to use ngrok for public access")
	exeSource = flag.String("payload", "", "URL or local path to the executable")
	useURL    = flag.Bool("url", true, "Set to true if exe is a URL, false if it is a local path")
)
