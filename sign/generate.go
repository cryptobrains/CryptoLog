package sign

// GenerateSignature() *dsa.PrivateKey

import (
	"crypto/rand"
	"crypto/dsa"
)

func GenerateSignature() *dsa.PrivateKey {

	params := new(dsa.Parameters)

	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		panic(err)
	}

	privatekey := new(dsa.PrivateKey)
	privatekey.PublicKey.Parameters = *params
	dsa.GenerateKey(privatekey, rand.Reader)

	return privatekey

}
