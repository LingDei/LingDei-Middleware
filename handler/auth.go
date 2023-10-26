package handler

import (
	"LingDei-Middleware/utils"
	"crypto/rsa"
)

var (
	PrivateKey *rsa.PrivateKey
)

func init() {
	var err error
	PrivateKey, err = utils.LoadPrivateKey()
	if err != nil {
		panic(err)
	}
}
