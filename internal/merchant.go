package internal

import (
	"context"
	merchantPb "github.com/rumsrami/givly-rpc-api/rpc/merchant"
)

// MerchantDirectory implements twirp MerchantService interface
type MerchantDirectory struct {}

// NewMerchantDirectory returns a new configured MerchantDirectory
func NewMerchantDirectory() (*MerchantDirectory) {
	return &MerchantDirectory{}
}

// ===================================================
// MerchantDirectory implementation of MerchantService
// ===================================================

// CreateMerchant creates a new merchant
func (m *MerchantDirectory) CreateMerchant(context.Context, *merchantPb.CreateMerchantRequest) (*merchantPb.Merchant, error) {
	/*
		Recieves storeemail and storename
		Checks for duplicate emails
		Hits endpoint for wallet
		Stores Merchant in db
		Returns back the new updated merchant information
	*/
	return nil, nil
}

// UpdateMerchant updates an existing merchant
func (m *MerchantDirectory) UpdateMerchant(context.Context, *merchantPb.UpdateMerchantRequest) (*merchantPb.Merchant, error) {
	/*
		Recieves merchantuid, storeemail, storename, wallet
		Saves to db only updated / provided fields
		Returns back the new updated merchant information
	*/
	return nil, nil
}

// VerifyBAccount checks if the Beneficiary has balance and returns balance
func (m *MerchantDirectory) VerifyBAccount(context.Context, *merchantPb.VerifyBAccountRequest) (*merchantPb.BAccountBalanceVerified, error){
	/*
		Revieves beneficiaryuid
		Hits endpoint for beneficiary wallet
		Returns false if beneficiary doesnt exit or balance amount
	*/
	return nil, nil
}