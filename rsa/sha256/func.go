package sha256

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// sha256 签名
func Sign(privateKey []byte, data string) (error, string) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return errors.New("private key error"), ""
	}
	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err, ""
	}
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	signedData, err := rsa.SignPKCS1v15(rand.Reader, private.(*rsa.PrivateKey), crypto.SHA256, hashed)
	return nil, base64.StdEncoding.EncodeToString(signedData)
}

// sha256 验证签名
func Verify(publicKey []byte, sign, data string) error {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("private key error")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	s, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed, s)
}
