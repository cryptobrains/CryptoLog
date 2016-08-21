package message

// MakeMessage(logMessage *LogMessage) *Message

import (
	globs "github.com/cryptobrains/CryptoLog/globs"
	db "github.com/cryptobrains/CryptoLog/db"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

func MakeMessage(logMessage *LogMessage) *Message {

	msg := &Message{}

	previousLogId := globs.GetLastLogId()

	logMessage.PreviousLogId = previousLogId

	timeCount := strings.SplitN(previousLogId, "|", 3)
	lastTime, lastCount := timeCount[1], timeCount[2]

	currentTime := time.Now().Truncate(time.Millisecond).String()
	currentTime = currentTime[:len(currentTime)-10]

	var currentCount int

	if currentTime != lastTime { currentCount = 0 } else { currentCount, _ := strconv.Atoi(lastCount); currentCount++ }

	id := "secure_log|" + currentTime + "|" + strconv.Itoa(currentCount)

	msg.Id = id
	msg.Payload = *logMessage

	globs.SetLastLogId(id)

	return msg
}

func ReadMsgFromDb(id string) (*Message, error) {

	s, _, coll := db.InitMongo("messages")
	defer s.Close()

	msg := &Message{}

	err := coll.Find(bson.M{"id":id}).One(&msg)

	return msg, err
}

func WriteMsgToDb(msg *Message) error {

	s, _, coll := db.InitMongo("messages")
	defer s.Close()

	if err := coll.Insert(*msg); err != nil {
		return err
	} else {
		return nil
	}
}
