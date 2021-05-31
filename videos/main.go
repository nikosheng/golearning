package main

import (
	"fmt"
)

func main() {
	videos := GetVideos()

	for i, _ := range videos {
		fmt.Println(videos[i])
		videos[i].Description = videos[i].Description + "\n Remember Like and Subscribe!"
	}

	SaveVideos(videos)
}
