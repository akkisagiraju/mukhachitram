package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	apiUrl := "https://ws.audioscrobbler.com/2.0/"

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(pwd)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}

	}

	// get pwd
	// loop through all the available directors inside pwd
	// make a get request to lastFM api with the directory name
	// download the album art that comes back in the response
	// save the image inside the corresponding directory

}
