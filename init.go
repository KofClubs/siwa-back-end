package main

import (
	"github.com/KofClubs/siwa-back-end/rsa2"
	"github.com/KofClubs/siwa-back-end/storage"
	"github.com/MonteCarloClub/log"
)

func init() {
	err := rsa2.GenerateKey()
	if err != nil {
		log.Error("fail to generate rsa2 key", "err", err)
		return
	}
	log.Info("rsa2 key generated")

	storage.InitClient()
	log.Info("storage client inited")
}
