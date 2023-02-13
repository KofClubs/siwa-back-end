package rsa2

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/KofClubs/siwa-back-end/config"
	"github.com/MonteCarloClub/log"
)

func GenerateKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Error("fail to generate rsa2 key", "err", err)
		return err
	}
	config.Rsa2PrivateKey, config.Rsa2PublicKey = privateKey, &privateKey.PublicKey
	privateKeyFile, err := os.Create(config.Rsa2PrivateKeyFilepath)
	if err != nil {
		log.Error(fmt.Sprintf("fail to create %v", config.Rsa2PrivateKeyFilepath), "err", err)
		return err
	}
	err = pem.Encode(privateKeyFile, &pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(config.Rsa2PrivateKey),
	})
	if err != nil {
		log.Error("fail to encode rsa2 private key", "err", err)
		return err
	}
	publicKeyFile, err := os.Create(config.Rsa2PublicFilepath)
	if err != nil {
		log.Error(fmt.Sprintf("fail to create %v", config.Rsa2PublicFilepath), "err", err)
		return err
	}
	err = pem.Encode(publicKeyFile, &pem.Block{
		Bytes: x509.MarshalPKCS1PublicKey(config.Rsa2PublicKey),
	})
	if err != nil {
		log.Error("fail to encode rsa2 public key", "err", err)
		return err
	}
	data, err := os.ReadFile(config.Rsa2PublicFilepath)
	if err != nil {
		log.Error("fail to read rsa2 public key file", "err", err)
		return err
	}
	config.Rsa2PublicKeyPem = string(data)
	log.Info("rsa2 key generated",
		"private key filepath", config.Rsa2PrivateKeyFilepath,
		"public key filepath", config.Rsa2PublicFilepath,
	)
	return nil
}

func Sign(message []byte) (string, error) {
	hash := sha512.New()
	hash.Write(bytes.NewBuffer(message).Bytes())
	sum := hash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, config.Rsa2PrivateKey, crypto.SHA512, sum)
	if err != nil {
		log.Error("fail to sign", "message", string(message), "err", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}
