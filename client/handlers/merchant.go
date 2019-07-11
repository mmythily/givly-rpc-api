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
}

// NewMerchantHandler returns a merchant subrouter
func NewMerchantHandler(addr string) MerchantHandler {
	merchantClient := merchantPb.NewMerchantServiceProtobufClient(addr, &http.Client{})

	return MerchantHandler{
		Client: merchantClient,
	}
}

// Route Mounts the merchant handlers on Router
func (m MerchantHandler) Route(r *mux.Router) {
	r.HandleFunc("/postme", m.handleCreateMerchant()).Methods("POST")
}

// handleCreateMerchant handles creating a new merchant
func (m MerchantHandler) handleCreateMerchant() http.HandlerFunc {
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
