package alive

// Serve()

import (
	ipc "github.com/cryptobrains/CryptoLog/ipc"
	"log"
)

type AlivePipe struct {
	*ipc.Pipe
}

const TIMEOUT int = 500; // waiting responses from pipe timeout
//const PATH string = "/tmp/securelog_ipc/alive_query/"
var PATH string

var canWork	chan bool

func Serve(path string) {
	PATH = path
	canWork = make(chan bool)
	pipe := &AlivePipe{
		ipc.NewPipe(PATH, TIMEOUT),
	}
	go pipe.serve()
	<- canWork
}

func (pipe *AlivePipe) waitForAliveQuery(){

	response := &ipc.IPCMessage{
		From: pipe.ID,
		Message: "TRUE",
	}

	go func(){
		for {
			inbox, err := pipe.ReceiveQuery()

			if err != nil {
				log.Println("Wrong message ", err)
				continue
			}

			log.Printf("Query is received from %s to %s, content: %s\n", inbox.From, inbox.To, inbox.Message.(string))
			query := inbox.Message.(string)

			if query != "ALIVE_QUERY" {
				log.Println("Unknown query: ", query)
				continue
			}

			response.To = inbox.From

			err = response.SendResponse()
			if err != nil { panic(err) }
			log.Printf("Response is sent from %s to %s, content: %s", response.From, response.To, response.Message.(string))
		}
	}()
}

func (pipe *AlivePipe) sendAliveQuery(pipename string) bool {

	if pipe.ID == pipename { return false }

	var response *ipc.IPCMessage = nil

	msg := &ipc.IPCMessage{
		From: pipe.ID,
		To: pipename,
		Message: "ALIVE_QUERY",
	}

	err := msg.SendQuery()
	if err != nil { panic(err) }

	log.Printf("Query is sent from %s to %s\n", msg.From, msg.To)

	response, _ = pipe.ReceiveResponse()

	if response == nil { return false }

	log.Println("Received response from %s to %s, content: %s", response.From, response.To, response.Message.(string))

	return (err == nil) && (response.From == msg.To) && (response.Message.(string) == "TRUE")
}

func (pipe *AlivePipe) serve() {
	files, err := pipe.GetAllWaitingQueries()

	if err != nil { panic(err) }

	for _, pipename := range(files) {

		if pipename == pipe.ID  { continue }

		if pipe.sendAliveQuery(pipename) {
			ipc.Remove(pipe.ID)
			panic("There is a server already running")
		} else {
			ipc.Remove(pipename)
		}
	}

	//if no running processes, run own
	canWork <- true
	pipe.waitForAliveQuery()
}
