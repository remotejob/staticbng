package mcontents

import (
	"fmt"
	"github.com/remotejob/comutils/gen"
	"github.com/remotejob/staticbng/mgenerator/dbgetall"
	"gopkg.in/mgo.v2"
	"testing"
)

func TestGenerate(t *testing.T) {

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	allrecords := dbgetall.GetAll(*dbsession)
	
	wordNum := gen.Random(1000,2000)
	
	mtext := Generate(wordNum, allrecords)

	fmt.Println(mtext)

}
