Package sign is about creating new DSA and converting it to a struct of strings to make it possible to keep DSA in Mongo which doesn't support big.Int format (and back)

func GenerateSignature() \*dsa.PrivateKey				// generates new DSA
func SignToStr(publickey \*dsa.PublicKey, r, s \*big.Int) \*SignStr	// converts DSA to a struct of strings
func StrToSign(signStr \*SignStr) \*Sign				// converts a struct of strings back to DSA
