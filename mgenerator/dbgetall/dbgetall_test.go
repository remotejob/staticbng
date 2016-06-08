package dbgetall

import (
	"gopkg.in/mgo.v2"
	"testing"
)

func TestGetAll(t *testing.T) {

	dbsession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	result := GetAll(*dbsession)

	if len(result) < 10 {

		t.Fatal("Cant be less 10")
	}

}
