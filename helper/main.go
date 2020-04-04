package helper

import (
	"log"
	"net/url"
)

// ResolveReference func
func ResolveReference(link string, reference string) string {
	u, err := url.Parse(reference)

	if err != nil {
		log.Fatal(err)
	}

	base, err := url.Parse(link)

	if err != nil {
		log.Fatal(err)
	}

	return base.ResolveReference(u).String()
}
