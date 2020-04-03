package helper

import (
	"log"
	"net/url"
)

func GetHost(link string) string {
	u, err := url.Parse(link)

	if err != nil {
		log.Println(err)
	}

	return u.Hostname()
}
