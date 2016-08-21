package msgserver

// New(timeLimit, msgLimit int) *MsgServer
// (msgServer *MsgServer) Start()
// (msgServer *MsgServer) SendMsg(msg *message.Message)

import (
	message "github.com/cryptobrains/CryptoLog/message"
	sign "github.com/cryptobrains/CryptoLog/sign"
	//db "github.com/cryptobrains/CryptoLog/db"
	globs "github.com/cryptobrains/CryptoLog/globs"
	"crypto/dsa"
	"time"
	"log"
)

type MsgServer struct {
	input	chan *message.Message
	TimeLimit	time.Duration
	MsgLimit	int
	StartTime	time.Time
	Messages	int
	privateKey	*dsa.PrivateKey
}

func New(timeLimit, msgLimit int) *MsgServer {

	msgServer := &MsgServer{}

	msgServer.input = make(chan *message.Message)
	msgServer.TimeLimit = time.Duration(timeLimit) * time.Millisecond
	msgServer.MsgLimit = msgLimit

	msgServer.reset()

	return msgServer
}

func (msgServer *MsgServer) reset(){

	msgServer.privateKey = sign.GenerateSignature()
	msgServer.Messages = 0
	msgServer.StartTime = time.Now()

}

func (msgServer *MsgServer) Start() {

//	db.ClearLogs()

	go func(){

		for msg := range(msgServer.input) {
			msgServer.Messages++
			duration := time.Since(msgServer.StartTime)

			if (duration >= msgServer.TimeLimit) || (msgServer.Messages >= msgServer.MsgLimit) {

				msgServer.reset()

			}

			msg.SignPayload(msgServer.privateKey)
			err := message.WriteMsgToDb(msg)
			if err != nil { log.Fatal("Writing to db error: ", err) }
		}
	}()
}

func (msgServer *MsgServer) SendMsg(msg *message.Message) {
	msgServer.input <- msg
}

func InitialPrev() string {
	return globs.InitialPrev
}

func GetLastLogId() string {
	return globs.GetLastLogId()
}
