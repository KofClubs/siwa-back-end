package handler

import (
	"encoding/json"
	"fmt"

	"github.com/KofClubs/siwa-back-end/rsa2"
	"github.com/MonteCarloClub/log"
	"github.com/MonteCarloClub/utils"
)

type SignRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	TradeNo    string `json:"trade_no"`
}

type SignResponse struct {
	Message string `json:"message"`
	Sign    string `json:"sign"`
}

func (signRequest *SignRequest) Handle() (*SignResponse, error) {
	if signRequest == nil {
		log.Error("/sign: nil request", "err", utils.NilPtrDerefErr)
		return nil, utils.NilPtrDerefErr
	}
	if signRequest.OutTradeNo == "" && signRequest.TradeNo == "" {
		err := fmt.Errorf("empty out_trade_no and trade_no")
		log.Error(fmt.Sprintf("/sign: %v", err.Error()), "err", err)
		return nil, err
	}
	message, err := json.Marshal(signRequest)
	if err != nil {
		log.Error("/sign: fail to marshal request", "err", err)
		return nil, err
	}
	sign, err := rsa2.Sign(message)
	if err != nil {
		log.Error("/sign: fail to sign request", "err", err)
		return nil, err
	}
	return &SignResponse{
		Message: string(message),
		Sign:    sign,
	}, nil
}
