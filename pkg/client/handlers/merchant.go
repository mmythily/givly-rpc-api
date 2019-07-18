package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/twitchtv/twirp"

	"github.com/gorilla/mux"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
)

// MerchantHandler handles merchant client requests
type MerchantHandler struct {
	Client merchantPb.MerchantService
	Router *mux.Router
}

// TwirpError defines a twirp Error
type TwirpError struct {
	Message string
	Status string
}

// HandleError handles errors
// Change this to Error Wrapper TODO
func HandleError(err error) []byte {
	//twirpErr := twirp.NewError()
	newError := TwirpError{
		Message: err.Error(),
		// Change to twirp code TODO
		Status: "",
	}
	errorToReturn, _ := json.Marshal(&newError)
	return errorToReturn
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
	m.Router.HandleFunc("/verifyAccount", m.handleGetRecipientBalance()).Methods("POST")
}

// handleCreateMerchant handles creating a new merchant
func (m MerchantHandler) handleCreateMerchant() http.HandlerFunc {
	// Code here gets run one time when instance starts
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		pbRequest := &merchantPb.CreateMerchantReq{
			StoreEmail: r.FormValue("storeEmail"),
			StoreName:  r.FormValue("storeName"),
		}

		pbResponse, err := m.Client.CreateMerchant(context.Background(), pbRequest)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			// Change status to match twirp error TODO
			w.WriteHeader(http.StatusNotFound)
			w.Write(HandleError(err))
		}
		
		fmt.Printf("pbres: %v+", pbResponse)
		response, _ := json.Marshal(pbResponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

// handleVerifyAccount verifies the account of a recipient
func (m MerchantHandler) handleGetRecipientBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
