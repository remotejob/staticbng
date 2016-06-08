package splitlink

import (
	"log"
	"net/url"
	"strings"
//	"github.com/remotejob/comutils/str"	
	//	"fmt"
)

func Split(link string) (string, string, []string) {

	var linkstr string
	var host string
	var rtitles []string
	var titles []string

	u, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	} else {
		host = u.Host
		linkstr = u.Host + u.Path
		rtitles = strings.Split(u.Path, "/")

		for i, title := range rtitles {

			if i > 0 {
				ctitle := cleantitle(title)
				titles = append(titles, ctitle)
			}
		}

		//		return u.Path

	}

	return host, linkstr, titles
}

func cleantitle(title string) string {

	ctitle := strings.Split(title, ".")[0]

	dashclean := strings.Split(ctitle, "-")

	if len(dashclean) > 0 {
		ctitle = ""
		for _, nodash := range dashclean {

			ctitle = ctitle + " " + nodash
		}
	}

	return ctitle
}
