package client

import (
	"github.com/rumsrami/givly-rpc-api/client/handlers"
)

// Client represents the twirp handlers
type Client struct {
	MerchantHandler    handlers.MerchantHandler
	TransactionHandler handlers.TransactionHandler
}

// New returns a new twirp Client
func New(addr string) *Client {

	merchantHandler := handlers.NewMerchantHandler(addr)
	transactionHandler := handlers.NewTransactionHandler(addr)

	return &Client{
		MerchantHandler:    merchantHandler,
		TransactionHandler: transactionHandler,
	}
}
