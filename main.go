package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

const (
	concurrency = 10
)

func main() {
	links := []string{"https://www.google.com/"}

	linkChan := make(chan string)
	nodeChan := make(chan *html.Node)

	go func() {
		for _, link := range links {
			linkChan <- link
		}
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			for link := range linkChan {
				nodes := fetch(link)

				for _, node := range nodes {
					go func(node *html.Node) {
						nodeChan <- node
					}(node)
				}
			}
		}()
	}

	for node := range nodeChan {
		fmt.Println(node)
	}
}

func fetch(url string) []*html.Node {
	client := &http.Client{
		Timeout: time.Duration(10 * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return nil
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)

	if err != nil {
		log.Println(err)
		return nil
	}

	return extract(node)
}

func extract(n *html.Node) []*html.Node {
	nodes := []*html.Node{}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			nodes = append(nodes, n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return nodes
}
