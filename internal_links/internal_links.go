package internal_links

import (
//	"fmt"
	"github.com/remotejob/staticbng/domains"
	"github.com/remotejob/staticbng/splitlink"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func Create(sitemapsObjs []domains.SitemapObj) []domains.LinkObj{

	//	fmt.Println("sitemapsObjs", sitemapsObjs[0])

	rand.Seed(int64(time.Now().Nanosecond()))
	records_num := len(sitemapsObjs)
	//	fmt.Println("records_num", records_num)
	numbers := rand.Perm(records_num)

	var internallinks []domains.LinkObj

	for i := 0; i < 10; i++ {

		_, _, titles := splitlink.Split(sitemapsObjs[numbers[i]].Loc)

		var urltitle string

		for _, title := range titles {

			urltitle = urltitle + " " + title

		}

//		fmt.Println(strings.TrimSpace(urltitle))

		urlstr := pathstr(sitemapsObjs[numbers[i]].Loc)

//		fmt.Println(urlstr)

		linkObjs := domains.LinkObj{Urlstr: urlstr, Title: strings.TrimSpace(urltitle)}

		internallinks = append(internallinks, linkObjs)

	}

return internallinks
}

func pathstr(urlstr string) string {

	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}

	return u.Path
}
