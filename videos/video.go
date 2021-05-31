package main

import (
	"encoding/json"
	"io/ioutil"
)

type (
	video struct {
		Id          string `json:"video_id"`
		Title       string `json:"video_title"`
		Description string `json:"video_description"`
		Imageurl    string `json:"video_imageurl"`
		Url         string `json:"video_url"`
	}
)

func GetVideos() (videos []video) {
	// Read File
	fileBytes, err := ioutil.ReadFile("./video.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func SaveVideos(videos []video) {
	fileBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos_saved.json", fileBytes, 0644)

	if err != nil {
		panic(err)
	}
}
