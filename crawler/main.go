package crawler

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/memochou1993/image-crawler/helper"
	"golang.org/x/net/html"
)

// Gallery struct
type Gallery struct {
	Links  []string
	Images []Image
}

// Image struct
type Image struct {
	Link string
	Node *html.Node
}

// Query func
func (g *Gallery) Query(r *http.Request, key string) {
	g.Links = strings.Split(strings.Replace(r.URL.Query().Get(key), " ", "", -1), ",")
}

// Fetch func
func (g *Gallery) Fetch() []string {
	nodeChan := make(chan Image)

	throttle := make(chan struct{}, 10)
	terminal := make(chan struct{}, 0)
	linkGroup := sync.WaitGroup{}
	nodeGroup := sync.WaitGroup{}

	linkGroup.Add(len(g.Links))

	for _, link := range g.Links {
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
			g.Images = append(g.Images, image)
		case <-terminal:
			break Loop
		}
	}

	return g.Format()
}

// Format func
func (g *Gallery) Format() []string {
	images := []string{}

	for _, image := range g.Images {
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
