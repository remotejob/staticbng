package create_stat_html

import (
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var htmlcont = `<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>`

func Create(file string) {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		d1 := []byte(htmlcont)
		err := ioutil.WriteFile(file, d1, 0644)
		check(err)

	}

}

func CreateIndex(dir string) {

	file := dir + "/index.html"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		d1 := []byte(htmlcont)
		err := ioutil.WriteFile(file, d1, 0644)
		check(err)

	}

}
