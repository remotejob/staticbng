package domains

import ()

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
}