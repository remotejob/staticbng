package create_stat_html

import (
	//	"fmt"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/comutils/str"
	"github.com/remotejob/staticbng/domains"
//	"github.com/remotejob/staticbng/mgenerator/dbgetall"
	"github.com/remotejob/staticbng/mgenerator/mcontents"
	"html/template"
	"os"
	"path"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Create(file string, allrecords []string, host string, titles []string, internallinks []domains.LinkObj) {

	if _, err := os.Stat(file); os.IsNotExist(err) {

		mod_title := mod_titles(host, titles)
		wordNum := gen.Random(1000, 2000)
		mtext := mcontents.Generate(wordNum, allrecords)

		f, err := os.Create(file)
		if err != nil {
			//    log.Println("create file: ", err)
			check(err)
			return
		}

		pageObj := domains.TeplatePage{Mcontents: mtext, Host: host, Titles: mod_title, Internallinks: internallinks}

		lp := path.Join("templates", "layout.html")

		t, err := template.ParseFiles(lp)
		check(err)

		err = t.Execute(f, pageObj)
		check(err)

	}

}

func CreateIndex(dir string, allrecords []string, host string, titles []string, internallinks []domains.LinkObj) {

	file := dir + "/index.html"

	if _, err := os.Stat(file); os.IsNotExist(err) {

		mod_title := mod_titles(host, titles)

		wordNum := gen.Random(1000, 2000)
		mtext := mcontents.Generate(wordNum, allrecords)

		f, err := os.Create(file)
		if err != nil {
			//    log.Println("create file: ", err)
			check(err)
			return
		}

		pageObj := domains.TeplatePage{Mcontents: mtext, Host: host, Titles: mod_title, Internallinks: internallinks}

		lp := path.Join("templates", "layout.html")

		t, err := template.ParseFiles(lp)
		check(err)

		err = t.Execute(f, pageObj)
		check(err)

	}

}

func mod_titles(host string, titles []string) []string {

	var title0 string = host
	var title1 string = host
	var title2 string = host

	var mod_title []string

	for i, tlt := range titles {

		if i == 0 {

			title0 = str.UpcaseInitial(tlt)
			mod_title = append(mod_title, title0)

		} else if i == 1 {

			title1 = str.UpcaseInitial(tlt)
			mod_title = append(mod_title, title1)

		} else if i == 2 {

			title2 = str.UpcaseInitial(tlt)
			mod_title = append(mod_title, title2)
		}

	}

	return mod_title

}
