Package message is for managing log messages - wrapping them up, signing etc.

func NewLogMessage(set of arguments depends on the structure of your log message) LogMessage	// creates log message, for using in app which sends messages to RabbitMQ
func MakeMessage(logMessage \*LogMessage) \*Message	// wraps log messages as payload to global messages

func (msg \*Message) SignPayload(privatekey \*dsa.PrivateKey) // signs payload by DSA
func (msg \*Message) VerifyPayload() bool			// verifies if ayload is properly signed
