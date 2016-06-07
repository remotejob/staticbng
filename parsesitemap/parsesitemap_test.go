package parsesitemap

import (
    "testing"
)

func TestParse(t *testing.T) {


	_,err := Parse("/home/juno/git/go_cv/version_00/maps/sitemap_127.0.0.1.xml")
	
	if err != nil {
		
		t.Fatal(err.Error())
	}
	
	

}

