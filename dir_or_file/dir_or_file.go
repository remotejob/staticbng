package dir_or_file

import (
	"strings"
)

func CheckifFile(link string) bool {

	if strings.HasSuffix(link, ".html") {

		return true
	}
	return false

}
