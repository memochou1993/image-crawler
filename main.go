package main

import "github.com/memochou1993/image-crawler/crawler"

func main() {
	links := []string{
		"https://www.104.com.tw/jobs/main/",
		"https://www.google.com/?1",
		"https://www.google.com/?2",
		"https://www.google.com/?3",
		"https://www.google.com/?4",
	}

	crawler.Initialize(links)
}
