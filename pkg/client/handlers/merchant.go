package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/golang/protobuf/proto"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
)

// mRPCProcessor represents the rpc function caller
type mRPCProcessor func(*http.Request, merchantPb.MerchantService) (proto.Message, error)

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
	m.Router.ServeHTTP(w, r)
}

// route Mounts the merchant handlers on Router
func (m MerchantHandler) route() {
	m.Router.HandleFunc("/create", m.respond(createMerchant, handleError)).Methods("POST")
	m.Router.HandleFunc("/verifyAccount", m.respond(verifyAccount, handleError)).Methods("POST")
}

// respond wraps the response with headers and logging
func (m *MerchantHandler) respond(process mRPCProcessor, format errFormater) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pbResponse, err := process(r, m.Client)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			// Change status to match twirp error TODO
			w.WriteHeader(http.StatusNotFound)
			w.Write(format(err))
		}
		// Marshall the response
		response, _ := json.Marshal(pbResponse)
		// Send back to Client
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

/*
	mRPCProcessors definitions
*/

/*
expects:
{
	storeEmail:
	storeName:
}
*/
// createMerchant handles creating a new merchant
func createMerchant(r *http.Request, pb merchantPb.MerchantService) (proto.Message, error) {
	// Parse incoming JSON
	r.ParseForm()
	// Create protobuf request
	pbRequest := &merchantPb.CreateMerchantReq{
		StoreEmail: r.FormValue("storeEmail"),
		StoreName:  r.FormValue("storeName"),
	}
	// Call RPC function and get protobuf response
	pbResponse, err := pb.CreateMerchant(context.Background(), pbRequest)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}

/*
expects:
{
	recipientCryptoId: "1234"
}
*/
// verifyAccount verifies the account of a recipient
func verifyAccount(r *http.Request, pb merchantPb.MerchantService) (proto.Message, error) {
	// Parse incoming JSON
	r.ParseForm()
	// Create protobuf request
	pbRequest := merchantPb.RecipientBalanceReq{
		RecipientCryptoId: r.FormValue("recipientCryptoId"),
	}
	// Call RPC function and get protobuf response
	pbResponse, err := pb.GetRecipientBalance(context.Background(), &pbRequest)
	if err != nil {
		return nil, err
	}
	return pbResponse, nil
}