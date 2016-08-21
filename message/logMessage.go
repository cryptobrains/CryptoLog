package message

// NewLogMessage(key0 int, key1 string) LogMessage

//for testing purposes

func NewLogMessage(content string) *LogMessage {

	newLogMsg := &LogMessage{
		Content: content,
	}

	return newLogMsg

}
