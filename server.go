package main

import (
	msgserver "github.com/cryptobrains/CryptoLog/msgserver"
	alive "github.com/cryptobrains/CryptoLog/alive"
	message "github.com/cryptobrains/CryptoLog/message"
	"github.com/streadway/amqp"
	"log"
	"encoding/json"
)

func Serve(queueName string, timeLimit, msgLimit int) {

	alive.Serve("/tmp/securelog_ipc/alive_query/") // check if one process is already running and if no start own "alive" process

	log.Println("CryptoLog has gone logging...")

	msgServer := msgserver.New(timeLimit, msgLimit)
	msgServer.Start()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil { panic(err) }
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil { panic(err) }
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, //name
		false, //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)

	for {
		d, ok, err := ch.Get(q.Name, false)
		if !ok {continue}
		if err != nil {log.Println(err)} else {
			logmsg, err := bodyToMsg(d.Body)
			if err != nil { log.Println(err) } else {
				msg := message.MakeMessage(logmsg)
				go msgServer.SendMsg(msg)
				d.Ack(false)
			}
		}
	}

}

func bodyToMsg(body []byte) (*message.LogMessage, error) {
	msg := message.LogMessage{}
	err := json.Unmarshal(body, &msg)
	return &msg, err
}

func main(){
	Serve("CryptoLog", 1000, 2000)
}
