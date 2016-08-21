Package msgserver is about serving log messages - it creates a server, starts it, accepts messages, signs them and writes to database


func New(timeLimit, msgLimit int) \*MsgServer			// creates a new server where timeLimit and msgLimit are the values
								// after achieving one of them new DSA is formed
func (msgServer \*MsgServer) Start()				// starts a server
func (msgServer \*MsgServer) SendMsg(msg \*message.Message)	// sends \*message.Message to a server
