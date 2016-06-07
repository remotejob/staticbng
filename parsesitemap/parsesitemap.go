package parsesitemap

import (
	"encoding/xml"
	//	"fmt"
	"io/ioutil"
	"time"
)

type SitemapObj struct {
	Changefreq    string
	Hoursduration float64
	Loc           string
	Lastmod       string
}
type Pages struct {
	//	Version string   `xml:"version,attr"`
	XMLName xml.Name `xml:"urlset"`
	XmlNS   string   `xml:"xmlns,attr"`
	//	XmlImageNS string   `xml:"xmlns:image,attr"`
	//	XmlNewsNS  string   `xml:"xmlns:news,attr"`
	Pages []*Page `xml:"url"`
}

type Page struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	Lastmod    string   `xml:"lastmod"`
	Changefreq string   `xml:"changefreq"`
	//	Name       string   `xml:"news:news>news:publication>news:name"`
	//	Language   string   `xml:"news:news>news:publication>news:language"`
	//	Title      string   `xml:"news:news>news:title"`
	//	Keywords   string   `xml:"news:news>news:keywords"`
	//	Image      string   `xml:"image:image>image:loc"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Parse(sitemapfile string) ([]SitemapObj,error) {

	dat, err := ioutil.ReadFile(sitemapfile)
	check(err)

	var sitemap Pages
	xml.Unmarshal(dat, &sitemap)

	var sitemapObjs []SitemapObj

	for _, page := range sitemap.Pages {

		layout := "2006-01-02T15:04:05+03:00"
		t, err := time.Parse(layout, page.Lastmod)
		check(err)
		duration := time.Since(t)

		sitemapObj := SitemapObj{page.Changefreq, duration.Hours(), page.Loc, page.Lastmod}

		sitemapObjs = append(sitemapObjs, sitemapObj)

	}

	return sitemapObjs,err

}
