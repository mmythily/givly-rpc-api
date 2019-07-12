package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	merchantPb "github.com/rumsrami/givly-rpc-api/rpc/merchant"
)

// MerchantHandler handles merchant client requests
type MerchantHandler struct {
	Client merchantPb.MerchantService
	Router *mux.Router
}

// NewMerchantHandler returns a merchant handler
func NewMerchantHandler(addr string, r *mux.Router) MerchantHandler {
	merchantClient := merchantPb.NewMerchantServiceProtobufClient(addr, &http.Client{})
	merchantHandler := MerchantHandler{
		Client: merchantClient,
		Router: r,
	}
	merchantHandler.route()
	return merchantHandler
}

// ServeHTTP implements Handler
func (m MerchantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Router.ServeHTTP(w,r)
}

// route Mounts the merchant handlers on Router
func (m MerchantHandler) route() {
	m.Router.HandleFunc("/create", m.handleCreateMerchant()).Methods("POST")
	m.Router.HandleFunc("/update", m.handleUpdateMerchant()).Methods("POST")
	m.Router.HandleFunc("/verifyAccount", m.handleVerifyAccount()).Methods("POST")
}

// handleCreateMerchant handles creating a new merchant
func (m MerchantHandler) handleCreateMerchant() http.HandlerFunc {
	// Code here gets run one time when instance starts
	return func(w http.ResponseWriter, r *http.Request) {
		pbRequest := &merchantPb.CreateMerchantRequest{
			Storeemail: "",
			Storename:  "",
			Wallet:     "",
		}

		pbResponse, err := m.Client.CreateMerchant(context.Background(), pbRequest)
		if err != nil {
			fmt.Println(err)
		}
		response, _ := json.Marshal(pbResponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// handleUpdateMerchant updates an existing merchant account
func (m MerchantHandler) handleUpdateMerchant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// handleVerifyAccount verifies the account of a recipient
func (m MerchantHandler) handleVerifyAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
