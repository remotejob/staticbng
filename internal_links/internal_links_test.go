package internal_links

import (
	"fmt"
	"github.com/remotejob/staticbng/parsesitemap"
	"testing"
)

func TestCreate(t *testing.T) {

	
	sitemapObjs, err := parsesitemap.Parse("/home/juno/git/go_cv/version_00/maps/sitemap_127.0.0.1.xml")
	if err != nil {
		fmt.Println(err.Error())
	}	

	Create(sitemapObjs)

}
