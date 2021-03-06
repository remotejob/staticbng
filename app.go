package main

import (
	"flag"
	"fmt"
	"github.com/remotejob/comutils/fls"
	"github.com/remotejob/staticbng/create_stat_html"
	"github.com/remotejob/staticbng/dir_or_file"
	"github.com/remotejob/staticbng/domains"
	"github.com/remotejob/staticbng/mgenerator/dbgetall"
	"github.com/remotejob/staticbng/parsesitemap"
	"github.com/remotejob/staticbng/splitlink"
	"github.com/remotejob/staticbng/internal_links"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"path/filepath"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	var sitemapdir string
	var webrootdir string
	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		sitemapdir = cfg.Dirs.Sitemapdir
		webrootdir = cfg.Dirs.Webrootdir		

	}

	searchDir := sitemapdir

	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	allrecords := dbgetall.GetAll(*dbsession)

	for _, sitemap_file := range fileList {
		fmt.Println(sitemap_file)

		sitemapObjs, err := parsesitemap.Parse(sitemap_file)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, sitemapObj := range sitemapObjs {

				host, linkpath, titles := splitlink.Split(sitemapObj.Loc)
				
				internallinks := internal_links.Create(sitemapObjs )				
				
				if dir_or_file.CheckifFile(linkpath) {

					fls.CreateDirForFile(webrootdir, linkpath)
					create_stat_html.Create(webrootdir+linkpath, allrecords, host, titles, internallinks)

				} else {
					fls.CreateDirForDir(webrootdir, linkpath)
					create_stat_html.CreateIndex(webrootdir+linkpath, allrecords, host, titles,internallinks)

				}

			}

		}
	}

}
