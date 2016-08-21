package ipc

// NewPipe(pathArg string, timeoutArg int) *Pipe
// (msg *IPCMessage) SendQuery() error
// (msg *IPCMessage) SendResponse() error
// (pipe *Pipe) ReceiveQuery() (*IPCMessage, error)
// (pipe *Pipe) ReceiveResponse() (*IPCMessage, error)
// Remove(pipename string)

import (
	utils "github.com/cryptobrains/CryptoLog/utils"
	"syscall"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"time"
	"log"
)

type Pipe struct {
	ID	string
}

type IPCMessage struct {
	From	string
	To	string
	Message	interface{}
}

var path string
var timeout	int

func NewPipe(pathArg string, timeoutArg int) *Pipe {
	timeout = timeoutArg
	if timeout <= 0 { panic("timeout less than or equal 0") }
	path = pathArg
	pipe := createPipe()
	return pipe
}

func makePipeDirs(){
	utils.Makedir(path + "q/")
	utils.Makedir(path + "a/")
}

func createPipe() *Pipe {

	var pipename string

	makePipeDirs()

	for {
		pipename = utils.CW(8)

		if !scanForSameName(pipename, "q/") && !scanForSameName(pipename, "a/") {
			break
		}
	}

	syscall.Mknod(path + "q/" + pipename, syscall.S_IFIFO|0666, 0); // making named pipe for queries to answer
	syscall.Mknod(path + "a/" + pipename, syscall.S_IFIFO|0666, 0); // making named pipe for getting answers for own queries

	return &Pipe{
		ID: pipename,
	}

}

func scanForSameName(pipename, subpath string) bool {

	if file, err := os.Stat(path + subpath + pipename); err == nil {
		if file.Name() == pipename {
			if file.Mode().String()[0] != 'p' {
				os.Remove(path + file.Name())
			} else { return true }
		}
	}

	return false
}

func (msg *IPCMessage) SendQuery() error {
	return send(msg, "q/")
}

func (msg *IPCMessage) SendResponse() error {
	return send(msg, "a/")
}

func send(msg *IPCMessage, subpath string) error {

	to := msg.To

	var pipename string

	pipename = path + subpath + to

	pipe, err := os.OpenFile(pipename, os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}

	defer pipe.Close()

	msgStr, err := json.Marshal(msg)

	if err != nil { return err }

	cmd := exec.Command("echo", string(msgStr))
	cmd.Stdout = pipe

	err = cmd.Run()

	if err != nil { return err }

/*	enc := json.NewEncoder(pipe)
	err = enc.Encode(*msg)*/

	cmd.Wait()

	return err
}

func (pipe *Pipe) ReceiveQuery() (*IPCMessage, error) {

	//var inboxBytes []byte

	inbox := IPCMessage{}

	pipename := path + "q/" + pipe.ID
	pipefile, err := os.Open(pipename)
	//pipefile, err := os.OpenFile(pipename, os.O_RDWR, 0660)

	if err != nil {
		panic(err)
	}

	defer pipefile.Close()

	/*cmd := exec.Command("cat")

	cmd.Stdin = pipefile

	inboxBytes, _ = cmd.Output()

	err = json.Unmarshal(inboxBytes, &inbox)*/

	decoder := json.NewDecoder(pipefile)
	err = decoder.Decode(&inbox)

	return &inbox, err

}

func (pipe *Pipe) ReceiveResponse() (*IPCMessage, error) {

	var err error

	inbox := IPCMessage{}

	pipename := path + "a/" + pipe.ID
	pipefile, err := os.OpenFile(pipename, os.O_RDWR, 0660)

	if err != nil {
		panic(err)
	}

	defer pipefile.Close()

	decoder := json.NewDecoder(pipefile)

	received := make(chan bool)

	go func(){
		err = decoder.Decode(&inbox)
		if err != nil { log.Println("Command run error: ", err) }
		received <- true
	}()

	select {
		case <-received:
//			log.Println(inbox)
		case <- time.After(time.Duration(timeout) * time.Millisecond):
//			log.Println("response timeout")
			return nil, errors.New("Receiving failed")
	}

	return &inbox, nil

}

func (pipe *Pipe) GetAllWaitingQueries() ([]string, error) {
	dir, err := os.Open(path + "q/")
	if err != nil {
		makePipeDirs()
		return []string{}, nil
	}
	defer func(){
		dir.Close()
	}()

	return dir.Readdirnames(0)
}

func Remove(pipename string) {
	os.Remove(path + "q/" + pipename)
	os.Remove(path + "a/" + pipename)
}
