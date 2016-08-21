To run a service just run from this folder 'go run server.go' or execute 'go build' and then use the executable got from it.

Example:

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/cryptobrains/CryptoLog/message"
	"strconv"
)

func main(){

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil { panic(err) }
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil { panic(err) }
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"secure_log", //name
		false, //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)

	if err != nil { panic(err) }

	for i := 0; i < 40; i++ {
		value := "I am a message number " + strconv.Itoa(i)

		msg := message.NewLogMessage(i, value)

		body, err := json.Marshal(msg)

		err = ch.Publish(
			"", //exchange
			q.Name, //routing key
			false, //mandatory
			false, //immediate
			amqp.Publishing {
				ContentType: "text/plain",
				Body: body,
			},
		)
		if err != nil { panic(err) }
	}
}
```
