package main

import (
	"flag"
	"fmt"
	"github.com/remotejob/staticbng/parsesitemap"
	"github.com/remotejob/staticbng/splitlink"
	"github.com/remotejob/staticbng/dir_or_file"
	"github.com/remotejob/staticbng/create_stat_html"
	"github.com/remotejob/comutils/fls"	
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}
	sitemapfile := "/home/juno/git/go_cv/version_00/maps/sitemap_127.0.0.1.xml"
	rootdir :="/tmp/"
	sitemapObjs, err := parsesitemap.Parse(sitemapfile)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, sitemapObj := range sitemapObjs {

			fmt.Println(splitlink.Split(sitemapObj.Loc))
			linkpath :=splitlink.Split(sitemapObj.Loc)
			if dir_or_file.CheckifFile(linkpath) {
				
				fls.CreateDirForFile(rootdir,linkpath)
				create_stat_html.Create(rootdir+linkpath)				
				
				
			} else {
				fls.CreateDirForDir(rootdir,linkpath)
				create_stat_html.CreateIndex(rootdir+linkpath)
				
			}
			
			
			

		}

	}

}
