package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	apiUrl := "https://ws.audioscrobbler.com/2.0/"

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get pwd
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// get all directories inside the pwd
	files, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	// loop through all the available directors inside pwd
	for _, file := range files {
		if file.IsDir() && file.Name() != ".git" {
			fmt.Println(file.Name())
			// make a get request to lastFM api with the directory name
			us := apiUrl + "?method=album.search&album=" + url.QueryEscape(file.Name()) + "&api_key=" + os.Getenv("APIKEY") + "&limit=1&format=json"
			resp, err := http.Get(us)
			if err != nil {
				log.Fatal(err)
			}
			body, err := io.ReadAll(resp.Body)
			defer resp.Body.Close()

			j := make(map[string]interface{})
			json.Unmarshal(body, &j)

			// download the album art that comes back in the response
			// save the image inside the corresponding directory

			fmt.Println(j)
		}

	}

}
