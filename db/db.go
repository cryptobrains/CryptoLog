package db

// InitMongo(collname string) (*mgo.Session, *mgo.Database, *mgo.Collection)
// ClearLogs()

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const DBPREFIX string = "securelog"

func InitMongo(collname string) (*mgo.Session, *mgo.Database, *mgo.Collection) {

	if collname == "" {panic("collection name is empty")}

	session,err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(DBPREFIX)
	coll := db.C(collname)
	return session, db, coll
}

func ClearLogs() {

	clearCollection("messages")
	clearCollection("globs")

}

func clearCollection(collection string) {

	s, _, coll := InitMongo(collection)
	defer s.Close()

	_, err := coll.RemoveAll(bson.M{})
	if err != nil {
		panic(err)
	}
}
