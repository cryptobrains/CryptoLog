db package is used to work with db

Endpoints:
func InitMongo(collname string) (\*mgo.Session, \*mgo.Database, \*mgo.Collection)
func ClearLogs()

Usage:

	import "github.com/cryptobrains/CryptoLog/db"

	//this opens the collection to work with
	session, db, coll := db.InitMongo("myDB", "myCollection")
	defer session.Close()

	//this clears all logs if needed
	db.ClearLogs()
