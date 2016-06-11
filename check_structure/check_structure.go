package main

import (
//		"fmt"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/staticbng/check_structure/scaner"
	"github.com/remotejob/staticbng/create_stat_html"
	"github.com/remotejob/staticbng/domains"
	"github.com/remotejob/staticbng/mgenerator/dbgetall"
	"github.com/remotejob/staticbng/mgenerator/mcontents"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

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

	var dirs_to_check []string

	var site string
	for _, sitemap_file := range fileList {

		site_spl0 := strings.Split(sitemap_file, "/")
		site_spl1 := site_spl0[len(site_spl0)-1]
		site_spl2 := strings.Split(site_spl1, "_")[1]
		site = strings.Replace(site_spl2, ".xml", "", 1)

		dirs_to_check = append(dirs_to_check, webrootdir+"/"+site)

	}

	for _, dir_to_check := range dirs_to_check {
		
//		fmt.Print("dir_to_check",dir_to_check)
		noIndexListDirs := scaner.Scan(dir_to_check)

		for _, noIndexListDir := range noIndexListDirs {

			wordNum := gen.Random(1000, 2000)
			mtext := mcontents.Generate(wordNum, allrecords)
			titles := []string{"index", "index", "index"}
			create_stat_html.CreateIndex(noIndexListDir, mtext, site, titles)
		}

	}

}
