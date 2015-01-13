package webctx

import (
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey []byte
	signKey   []byte
)

func SignToken(token *jwt.Token) (string, error) {
	return token.SignedString(signKey)
}

func LoadSecureKeys(privateKeyPath, publicKeyPath string) (err error) {
	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("Error reading private key")
	}
	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("Error reading public key")
	}
	return nil
}
