package domains

import (
"encoding/xml"
)

type LinkObj struct {
	
	Urlstr string
	Title string
}


type ServerConfig struct {
	Dirs struct {
		Sitemapdir string
		Webrootdir string
	}
}


type TeplatePage struct {
	Mcontents string
	Host      string
	Titles    []string
	Internallinks []LinkObj 
}

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