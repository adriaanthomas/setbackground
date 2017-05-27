package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/reujab/wallpaper"
)

func main() {
	fmt.Printf("Setting background..\n")

	js, err := loadImageData()
	if err != nil {
		log.Fatal(err)
	}

	img, err := decodeImageData(js)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://www.bing.com" + img.URL
	fmt.Printf("Image URL: %s\n", url)

	err = setBackground(url)
	if err != nil {
		log.Fatal(err)
	}
}

type message struct {
	Images []image `json:"images"`
}

type image struct {
	URL string `json:"url"`
}

func loadImageData() (js []byte, err error) {
	resp, e := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&pid=hp")
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()

	js, err = ioutil.ReadAll(resp.Body)
	return
}

func decodeImageData(js []byte) (image image, err error) {
	var msg message
	err = json.Unmarshal(js, &msg)
	if err != nil {
		return
	}

	image = msg.Images[0]
	return
}

func setBackground(url string) (err error) {
	err = wallpaper.SetFromURL(url)
	return
}
