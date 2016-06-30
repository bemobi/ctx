package ctx

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func SignToken(token *jwt.Token) (string, error) {
	return token.SignedString(signKey)
}

func LoadSecureKeys(privateKeyPath, publicKeyPath string) (err error) {
	bytesSignKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return fmt.Errorf("Error reading private key")
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytesSignKey)
	if err != nil {
		return fmt.Errorf("error parsing RSA private key: %s", err)
	}

	bytesVerifyKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("Error reading public key")
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(bytesVerifyKey)
	if err != nil {
		return fmt.Errorf("error parsing RSA private key: %s", err)
	}

	if err != nil {
		return fmt.Errorf("error parsing RSA public key: %s", err)
	}
	return nil
}
