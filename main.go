package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	var apiKey string

	flag.StringVar(&apiKey, "apikey", "", "Enter your LastFM API key")

	flag.Parse()

	if len(apiKey) == 0 {
		fmt.Printf("API key cannot be empty. \n")
		fmt.Printf("Usage: ./mukhachitram -apikey yourapikey \n")
		os.Exit(2)
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() && file.Name()[0] != '.' {
			// add an error handler if the length results is 0
			imageSlice := fetchAlbumDetails(file.Name(), apiKey).Results.Albummatches.Album[0].Image
			imgUrl := imageSlice[len(imageSlice)-1].Text
			downloadAndSaveImage(imgUrl, file.Name())
		}
	}

}

func fetchAlbumDetails(albumName string, apiKey string) AlbumResults {
	apiUrl := "https://ws.audioscrobbler.com/2.0/"
	us := apiUrl + "?method=album.search&album=" + url.QueryEscape(albumName) + "&api_key=" + apiKey + "&limit=1&format=json"

	resp, err := http.Get(us)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	j := AlbumResults{}
	json.Unmarshal(body, &j)

	return j
}

func downloadAndSaveImage(imgUrl string, dirName string) {
	response, err := http.Get(imgUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	imgFile, err := os.Create(filepath.Join(dirName, filepath.Base("cover.jpg")))
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()

	_, err = io.Copy(imgFile, response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}

	fmt.Printf("Saved album art of " + dirName + "\n")
}
