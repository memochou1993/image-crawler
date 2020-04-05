package crawler

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/memochou1993/image-crawler/helper"
	"golang.org/x/net/html"
)

// Collection struct
type Collection struct {
	Images []Image
}

// Image struct
type Image struct {
	Link string
	Node *html.Node
}

// Fetch func
func (c *Collection) Fetch(links []string) []string {
	throttle := make(chan struct{}, 10)
	terminal := make(chan struct{}, 0)
	linkGroup := sync.WaitGroup{}
	nodeGroup := sync.WaitGroup{}

	nodeChan := make(chan Image)

	linkGroup.Add(len(links))

	for _, link := range links {
		throttle <- struct{}{}

		go func(link string) {
			nodes := fetch(link)

			nodeGroup.Add(len(nodes))

			for _, node := range nodes {
				go func(node *html.Node) {
					nodeChan <- Image{
						Link: link,
						Node: node,
					}

					nodeGroup.Done()
				}(node)

				log.Println("sent node")
			}

			<-throttle

			linkGroup.Done()
		}(link)
	}

	go func() {
		linkGroup.Wait()
		nodeGroup.Wait()
		close(terminal)
	}()

Loop:
	for {
		select {
		case image := <-nodeChan:
			log.Println("received image")
			c.Images = append(c.Images, image)
		case <-terminal:
			break Loop
		}
	}

	return c.format()
}

func (c *Collection) format() []string {
	images := []string{}

	for _, image := range c.Images {
		for _, a := range image.Node.Attr {
			if a.Key == "src" && a.Val != "" {
				images = append(images, helper.ResolveReference(image.Link, a.Val))
			}
		}
	}

	return images
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
