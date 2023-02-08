package handler

import (
	"context"
	"encoding/json"

	"github.com/KofClubs/siwa-back-end/config"
	"github.com/MonteCarloClub/log"
	"github.com/MonteCarloClub/utils"
)

type UploadRequest struct {
	Message string `json:"message"`
	Sign    string `json:"sign"`
}

type UploadResponse struct {
	Key string `json:"key"`
}

func (uploadRequest *UploadRequest) Handle() (*UploadResponse, error) {
	if uploadRequest == nil {
		log.Error("/upload: nil request", "err", utils.NilPtrDerefErr)
		return nil, utils.NilPtrDerefErr
	}
	key := uploadRequest.Sign[:config.KeyLength]
	valueBytes, err := json.Marshal(uploadRequest)
	if err != nil {
		log.Error("/upload: fail to marshal request", "err", err)
		return nil, err
	}
	value := string(valueBytes)
	go func(key string, value string) {
		config.RedisClient.Set(context.Background(), key, value, 0)
	}(key, value)
	log.Info("/upload: asynchronous upload started", "key", key, "value", value)
	return &UploadResponse{
		Key: uploadRequest.Sign[:config.KeyLength],
	}, nil
}
