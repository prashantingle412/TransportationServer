package DbConfig
import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
)

var Collection *mgo.Collection

func SetCollection(dbName string, collectionName string) *mgo.Collection {
	var Collection *mgo.Collection
	var sess *mgo.Session
	if sess == nil {
		log.Println("Not connected... Connecting to Mongo")
		sess = GetConnected()
	}
	Collection = sess.DB(dbName).C(collectionName)
	return Collection
}

func GetConnected() *mgo.Session {
	var sess *mgo.Session
	dialInfo, err := mgo.ParseURL("mongodb://localhost:27017")
	dialInfo.Direct = true
	dialInfo.FailFast = true
	dialInfo.Database = "transportation_db"
	dialInfo.Username = "root"
	dialInfo.Password = "tiger"
	sess, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Println("Can't connect to mongo, go error %v\n", err)
		panic(err)
	} else {
		return sess
		defer sess.Close()
	}
	return sess
}