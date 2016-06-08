package dbgetall

import (
//	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

func GetAll(dbsession mgo.Session) []string {

	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("cv_employers").C("employers")

	var retresult []string

	var result []struct {
		Description string `bson:"description"`
	}

	err := c.Find(nil).Select(bson.M{"description": 1}).All(&result)
	if err != nil {
		// handle error
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	records_num := len(result)
	numbers := rand.Perm(records_num)
//	fmt.Println(numbers)

	for _, v := range numbers {
		//		fmt.Println(v.Description)
		retresult = append(retresult, result[v].Description)
	}

	return retresult
}
