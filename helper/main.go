package helper

import (
	"log"
	"net/url"
	"strings"
)

// ResolveReference func
func ResolveReference(link string, reference string) string {
	u, err := url.Parse(strings.TrimSpace(reference))

	if err != nil {
		log.Println(err)
	}

	base, err := url.Parse(strings.TrimSpace(link))

	if err != nil {
		log.Println(err)
	}

	return base.ResolveReference(u).String()
}
