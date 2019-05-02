package crypto


import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"github.com/pkg/errors"
	jwt "github.com/dgrijalva/jwt-go"
)


func LoadPrivBytes(keyfile string) ([]byte, error) {
	bKeyFile, err := ioutil.ReadFile(keyfile)
	if err != nil {
		err = errors.Wrap(err, "failed to read keyfile in LoadPrivRSA at location" + keyfile)
		return nil, err
	}
        block , _ := pem.Decode(bKeyFile)
	var bKey []byte = block.Bytes
	return bKey, err
}


func LoadPrivRSA(keyfile string) (*rsa.PrivateKey, error) {
	bKeyFile, err := ioutil.ReadFile(keyfile)
	if err != nil {
		err = errors.Wrap(err, "failed to read keyfile in LoadPrivRSA at location" + keyfile)
		return nil, err
	}
	rsaPriv, err := jwt.ParseRSAPrivateKeyFromPEM(bKeyFile)
	if err != nil {
		err = errors.Wrap(err, "failed to parsa RSA from key file bytes in LoadPrivRSA")
		return nil, err
	}
	return rsaPriv, err
}


func LoadPubRSA(certfile string) (*rsa.PublicKey, error) {
	var rsaPubKey *rsa.PublicKey
        bCertFile, err := ioutil.ReadFile(certfile)
	if err != nil {
		err = errors.Wrap(err, "failed to read certfile in LoadPubRSA at location" + certfile)
		return nil, err
	}
	block, _ := pem.Decode(bCertFile)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		err = errors.Wrap(err, "failed to parse certificate in LoadPubRSA")
		return nil, err
	}
	rsaPubKey = cert.PublicKey.(*rsa.PublicKey)
	return rsaPubKey, err
}

