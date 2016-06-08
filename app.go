package main

import (
	"flag"
	"fmt"
	"github.com/remotejob/comutils/fls"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/staticbng/create_stat_html"
	"github.com/remotejob/staticbng/dir_or_file"
	"github.com/remotejob/staticbng/mgenerator/dbgetall"
	"github.com/remotejob/staticbng/mgenerator/mcontents"
	"github.com/remotejob/staticbng/parsesitemap"
	"github.com/remotejob/staticbng/splitlink"
	"gopkg.in/mgo.v2"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	allrecords := dbgetall.GetAll(*dbsession)

	sitemapfile := "/home/juno/git/go_cv/version_00/maps/sitemap_127.0.0.1.xml"
	rootdir := "/tmp/"
	sitemapObjs, err := parsesitemap.Parse(sitemapfile)

	var mtext string
	var wordNum int

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, sitemapObj := range sitemapObjs {

//			fmt.Println(splitlink.Split(sitemapObj.Loc))
			host,linkpath,titles := splitlink.Split(sitemapObj.Loc)
			fmt.Println(linkpath)
			if dir_or_file.CheckifFile(linkpath) {

				fls.CreateDirForFile(rootdir, linkpath)
				wordNum = gen.Random(1000, 2000)
				mtext = mcontents.Generate(wordNum, allrecords)
				create_stat_html.Create(rootdir + linkpath,mtext,host,titles)

			} else {
				fls.CreateDirForDir(rootdir, linkpath)

				wordNum = gen.Random(1000, 2000)
				mtext = mcontents.Generate(wordNum, allrecords)
				create_stat_html.CreateIndex(rootdir + linkpath,mtext,host,titles)

			}

		}

	}

}
