package sign

// SignToStr(publickey *dsa.PublicKey, r, s *big.Int) *SignStr
// StrToSign(signStr *SignStr) *Sign

import (
	"crypto/dsa"
	"math/big"
	utils "github.com/cryptobrains/CryptoLog/utils"
)

type ParametersStr struct {
	P	string
	Q	string
	G	string
}

type PublicKeyStr struct {
	Y	string
	Parameters	*ParametersStr
}

type SignStr struct {
	PublicKey	*PublicKeyStr
	R	string
	S	string
}

type Sign struct {
	PublicKey	*dsa.PublicKey
	R	*big.Int
	S	*big.Int
}

func SignToStr(publickey *dsa.PublicKey, r, s *big.Int) *SignStr {

	signStr := SignStr{}
	publicKeyStr := PublicKeyStr{}
	paramStr := ParametersStr{}

	publicKeyStr.Y = publickey.Y.String()

	parameters := publickey.Parameters

	paramStr.P = parameters.P.String()
	paramStr.Q = parameters.Q.String()
	paramStr.G = parameters.G.String()

	publicKeyStr.Parameters = &paramStr

	signStr.PublicKey = &publicKeyStr

	signStr.R = r.String()
	signStr.S = s.String()

	return &signStr

}

func StrToSign(signStr *SignStr) *Sign {

	publickey := &dsa.PublicKey{}

	publickeyStr := signStr.PublicKey
	if publickeyStr == nil { return &Sign{} }

	paramStr := publickeyStr.Parameters

	publickey.Y = utils.StrToBigInt(publickeyStr.Y)

	publickey.Parameters = dsa.Parameters{
		P: utils.StrToBigInt(paramStr.P),
		Q: utils.StrToBigInt(paramStr.Q),
		G: utils.StrToBigInt(paramStr.G),
	}

	r := utils.StrToBigInt(signStr.R)
	s := utils.StrToBigInt(signStr.S)

	return &Sign{
		PublicKey: publickey,
		R: r,
		S: s,
	}

}
