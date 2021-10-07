package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube"
)

func downloader(url string) {
	videoID := url
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
	println("Downloaded /video.mp4")
}

func main() {
	var text string
	fmt.Println("Enter the url to downloaded: ")
	fmt.Scanf("%s", &text)
	fmt.Println(text)
	url := text
	downloader(url)
}
