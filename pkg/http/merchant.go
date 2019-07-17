package http

import (
	"context"
	"errors"

	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	"github.com/rumsrami/givly-rpc-api/pkg/adding"
	"github.com/rumsrami/givly-rpc-api/pkg/listing"
)

// MerchantDirectory implements twirp MerchantService interface
type MerchantDirectory struct{
	A adding.Service
	L listing.Service
}

// NewMerchantDirectory returns a new configured MerchantDirectory
func NewMerchantDirectory(a adding.Service, l listing.Service) *MerchantDirectory {
	return &MerchantDirectory{
		A: a,
		L: l,
	}
}

// ===================================================
// MerchantDirectory implementation of MerchantService
// ===================================================

/*
	For Error handling:
	Standard error is created here from the db returned error
	After database function gets called check for the error
	Return a readable standard Go error to the client
	The client will Make this error Twirp and JSONs it
	Sends it back to Caller with the HTTP status
	{
		"Message": "twirp error not_found: Cannot find email"
		"Status" : ""
	}
*/

// CreateMerchant creates a new merchant
func (m *MerchantDirectory) CreateMerchant(ctx context.Context, req *merchantPb.CreateMerchantReq) (*merchantPb.Merchant, error) {
	//var mytimestamp timestamp.Timestamp
	/*
		Recieves storeemail and storename
		Checks for duplicate emails
		Hits endpoint for wallet
		Stores Merchant in db
		Returns back the new updated merchant information
	*/

	if req.StoreEmail == "error" {
		return nil, errors.New("Email not found")
	}
	return &merchantPb.Merchant{
		MerchantUuid: "1234",
		StoreEmail:  "r@r.com",
		StoreName:   "Kleb",
	}, nil
}


// GetRecipientBalance checks if the Recipient has balance and returns balance
func (m *MerchantDirectory) GetRecipientBalance(ctx context.Context, req *merchantPb.RecipientBalanceReq) (*merchantPb.RecipientBalance, error) {
	/*
		Revieves recipientuid
		Hits endpoint for recipient wallet
		Returns false if recipient doesnt exit or balance amount
	*/
	return nil, nil
}

