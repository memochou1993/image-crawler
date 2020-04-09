package helper

import (
	"log"
	"net/url"
	"strings"
)

// Scheme func
func Scheme(link string) string {
	u, err := url.Parse(strings.TrimSpace(link))

	if err != nil {
		log.Println(err)
	}

	return u.Scheme
}

// Hostname func
func Hostname(link string) string {
	u, err := url.Parse(strings.TrimSpace(link))

	if err != nil {
		log.Println(err)
	}

	return u.Hostname()
}

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
