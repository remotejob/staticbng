package mcontents

import (
    "testing"
    "gopkg.in/mgo.v2"
    "github.com/remotejob/staticbng/mgenerator/dbgetall" 
    "fmt"   
)

func TestGenerate(t *testing.T) {
	
		dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()
	
	allrecords := dbgetall.GetAll(*dbsession)
	
	mtext := Generate(allrecords)
	
	fmt.Println(mtext)	

}

