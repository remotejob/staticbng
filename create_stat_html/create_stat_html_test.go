package create_stat_html

import (
	"testing"
)

func TestCreate(t *testing.T) {

	Create("/tmp/test/test2/test.html")
	CreateIndex("/tmp/test3/test4")
}
