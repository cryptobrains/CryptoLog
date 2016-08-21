package message

// (msg *Message) SignPayload(privatekey *dsa.PrivateKey)
// (msg *Message) VerifyPayload() bool

import (
	sign "github.com/cryptobrains/CryptoLog/sign"
	"crypto/dsa"
	"crypto/sha1"
	"crypto/rand"
	"encoding/json"
	"io"
)

type LogMessage struct {
	PreviousLogId	string

	Content	string
}

type Message struct {
	Id	string
	Payload	LogMessage
	SignedPayload []byte
	Signature	sign.SignStr
}

func (msg *Message) SignPayload(privatekey *dsa.PrivateKey) {

	payload := msg.Payload
	payloadJson, err := json.Marshal(payload)

	h := sha1.New()
	io.WriteString(h, string(payloadJson))

	hashedPayload := h.Sum(nil)


	if err != nil {
		panic(err)
	}

	r, s, err := dsa.Sign(rand.Reader, privatekey, hashedPayload)

	msg.SignedPayload = hashedPayload

	signature := sign.SignToStr(&privatekey.PublicKey, r, s)

	msg.Signature = *signature

}

func (msg *Message) VerifyPayload() bool {

	signature := sign.StrToSign(&msg.Signature)
	if (signature == &sign.Sign{}) { return false }

	publickey, r, s := signature.PublicKey, signature.R, signature.S

	if publickey == nil { return false }

	return dsa.Verify(publickey, msg.SignedPayload, r, s)

}
