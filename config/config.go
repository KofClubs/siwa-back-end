package config

import (
	"crypto/rsa"

	"github.com/go-redis/redis/v8"
)

const (
	Rsa2PrivateKeyFilepath = "rsa2_private_key.pem"
	Rsa2PublicFilepath     = "rsa2_public_key.pem"
	KeyLength              = 8
	RedisClientAddr        = "localhost:6379"
)

var (
	Rsa2PrivateKey   *rsa.PrivateKey
	Rsa2PublicKey    *rsa.PublicKey
	Rsa2PublicKeyPem string
	RedisClient      *redis.Client
)
