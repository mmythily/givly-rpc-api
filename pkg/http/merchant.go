package http

import (
	"context"

	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	//"github.com/golang/protobuf/ptypes/timestamp"
)

// MerchantDirectory implements twirp MerchantService interface
type MerchantDirectory struct{}

// NewMerchantDirectory returns a new configured MerchantDirectory
func NewMerchantDirectory() *MerchantDirectory {
	return &MerchantDirectory{}
}

// ===================================================
// MerchantDirectory implementation of MerchantService
// ===================================================

// CreateMerchant creates a new merchant
func (m *MerchantDirectory) CreateMerchant(ctx context.Context, req *merchantPb.CreateMerchantRequest) (*merchantPb.Merchant, error) {
	//var mytimestamp timestamp.Timestamp
	/*
		Recieves storeemail and storename
		Checks for duplicate emails
		Hits endpoint for wallet
		Stores Merchant in db
		Returns back the new updated merchant information
	*/
	return &merchantPb.Merchant{
		Merchantuid: "1234",
		Storeemail:  "r@r.com",
		Storename:   "Kleb",
		Wallet:      "1234",
	}, nil
}

// UpdateMerchant updates an existing merchant
func (m *MerchantDirectory) UpdateMerchant(ctx context.Context, req *merchantPb.UpdateMerchantRequest) (*merchantPb.Merchant, error) {
	/*
		Recieves merchantuid, storeemail, storename, wallet
		Saves to db only updated / provided fields
		Returns back the new updated merchant information
	*/
	return nil, nil
}

// VerifyRTAccount checks if the Recipient has balance and returns balance
func (m *MerchantDirectory) VerifyRTAccount(ctx context.Context, req *merchantPb.VerifyRTAccountRequest) (*merchantPb.RTAccountBalanceVerified, error) {
	/*
		Revieves recipientuid
		Hits endpoint for recipient wallet
		Returns false if recipient doesnt exit or balance amount
	*/
	return nil, nil
}

