Package ipc provides methods for sending and receiving queries and responses for interprocess communication

func NewPipe(pathArg string, timeoutArg int) \*Pipe		// creates a new pipe
func (msg \*IPCMessage) SendQuery() error			// sends query to the pipe which ID must be assigned to msg.To
func (msg \*IPCMessage) SendResponse() error			// send response to the pipe which ID must be assigned to msg.To
func (pipe \*Pipe) ReceiveQuery() (\*IPCMessage, error)		// waits for queries (eternal)
func (pipe \*Pipe) ReceiveResponse() (\*IPCMessage, error)	// waits for response (for a number of milliseconds determined by pipe's timeout parameter)
func Remove(pipename string)					// removes pipe files from disk

Usage example:

import "github.com/cryptobrains/CryptoLog/ipc"

pipe := ipc.NewPipe("/working/path/somewhere/on/disk/", 3000) // timeout is timeout in milliseconds for whic query will wait a response

msg := &ipc.IPCMessage({
	From: pipe.ID,
	To: otherPipeID,
	Message: "SOME_INFO", // this can be serialized json
})

msg.SendQuery()
response, err := pipe.RecieveResponse()

.... other process ....

query, err := pipe.ReceiveQuery()

if err != nil {
	msg := ipc.IPCMessage{
		From: pipe.ID,
		To: query.From,
		Message: "any response",
	}
	msg.SendResponse()
}
