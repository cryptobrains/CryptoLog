package globs

// GetLastLogId() string
// SetLastLogId(id string)
// InitialPrev

import (
	db "github.com/cryptobrains/CryptoLog/db"
	"gopkg.in/mgo.v2/bson"
)

type Global struct {
	Id		string
	LastLogId	string
}

const InitialPrev string = "secure_log|beginning|0"

func GetLastLogId() string {

	s, _, coll := db.InitMongo("globs")
	defer s.Close()

	global := Global{}

	if err := coll.Find(bson.M{"id":"global"}).One(&global); err != nil {
		return InitialPrev
	}

	lastLogId := global.LastLogId

	return lastLogId
}

func SetLastLogId(id string) {

	s, _, coll := db.InitMongo("globs")
	defer s.Close()

	global := Global{}

	global.Id = "global"
	global.LastLogId = id

	if _, err := coll.Upsert(bson.M{"id":"global"}, global); err != nil {
		panic(err)
	}

}
