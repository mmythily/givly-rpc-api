package internal

import (
	merchantPb "github.com/rumsrami/givly-rpc-api/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/rpc/transaction"
)

// NewMerchantDirectoryHandler returns a new Merchant handler
func NewMerchantDirectoryHandler() merchantPb.TwirpServer {
	mdServer := NewMerchantDirectory()
	twirpmdHandler := merchantPb.NewMerchantServiceServer(mdServer, nil)
	return twirpmdHandler
}

// NewTransactionDirectoryHandler returns a new Transaction handler
func NewTransactionDirectoryHandler() transactionPb.TwirpServer {
	tdServer := NewTransactionDirectory()
	twirptdHandler := transactionPb.NewTransactionServiceServer(tdServer, nil)
	return twirptdHandler
}