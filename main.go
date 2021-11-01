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

type AlbumResults struct {
	Results struct {
		Albummatches struct {
			Album []AlbumStruct `json:"album"`
		} `json:"albummatches"`
	} `json:"results"`
}

type AlbumStruct struct {
	Name   string        `json:"name"`
	Artist string        `json:"artist"`
	URL    string        `json:"url"`
	Image  []ImageStruct `json:"image"`
}

type ImageStruct struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get pwd
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// get all directories inside pwd
	files, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	// loop through all the available directors inside pwd
	for _, file := range files {
		if file.IsDir() && file.Name() != ".git" {
			var imgUrl string
			// add an error handler if the length results is 0
			for _, img := range fetchAlbumDetails(file.Name()).Results.Albummatches.Album[0].Image {
				if img.Size == "extralarge" {
					imgUrl = img.Text
				}
			}
			fmt.Println(imgUrl)
		}
	}
}

func fetchAlbumDetails(name string) AlbumResults {
	apiUrl := "https://ws.audioscrobbler.com/2.0/"
	us := apiUrl + "?method=album.search&album=" + url.QueryEscape(name) + "&api_key=" + os.Getenv("APIKEY") + "&limit=1&format=json"

	resp, err := http.Get(us)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	j := AlbumResults{}
	json.Unmarshal(body, &j)

	return j
}
