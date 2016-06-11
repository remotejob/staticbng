package scaner

import (
	//	"log"
//	"fmt"
	"os"
	"path/filepath"
)

func Scan(searchDir string) []string{

	noIndexListDirs := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() {

			if _, err := os.Stat(path + "/index.html"); os.IsNotExist(err) {
				noIndexListDirs = append(noIndexListDirs, path)
				// path/to/whatever does not exist
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}

  return noIndexListDirs
}
