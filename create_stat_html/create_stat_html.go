package create_stat_html

import (
	"io/ioutil"
	"os"
	"github.com/remotejob/comutils/str"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func PageHtml(mcontents string, host string, titles []string) string {

	var title0 string =host
	var title1 string = host
	var title2 string =host
	
	for i,tlt := range titles {
		
		if i == 0 {
			
			title0 = str.UpcaseInitial(tlt)			
			
		} else if i == 1 {
			
			title1 = str.UpcaseInitial(tlt)
						
		} else if i == 2 {
			
			title2 = str.UpcaseInitial(tlt)			
		}
		
		
	}
	
	

	var htmlcont = `<!DOCTYPE html>
<html>
<head>
<title>`+title2+`</title>
</head>
<body>

<h1>`+title0+`</h1>
<h2>`+title1+`</h2>

<p>`+title2+`</p>	` + mcontents + `

</body>
</html>`
	return htmlcont
}

func Create(file string, mcontents string, host string, titles []string) {

	if _, err := os.Stat(file); os.IsNotExist(err) {

		htmlcont := PageHtml(mcontents,host,titles)

		d1 := []byte(htmlcont)
		err := ioutil.WriteFile(file, d1, 0644)
		check(err)

	}

}

func CreateIndex(dir string, mcontents string, host string, titles []string) {

	htmlcont := PageHtml(mcontents,host,titles)

	file := dir + "/index.html"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		d1 := []byte(htmlcont)
		err := ioutil.WriteFile(file, d1, 0644)
		check(err)

	}

}
