package splitlink

import (
	"log"
	"net/url"
)

func Split(link string) string {

	var linkstr string

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	} else {
		linkstr = u.Host+u.Path
		//		return u.Path

	}

	return linkstr
}
