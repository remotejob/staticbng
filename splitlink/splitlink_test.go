package splitlink

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {

	host, path, titles := Split("http://127.0.0.1/blog")
	fmt.Println(host, path, titles)

	host, path, titles = Split("http://127.0.0.1/blog/ruby/convert-sql-query-to-active-record.html")
	fmt.Println(host, path, titles)
}
