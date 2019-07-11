package http

import (
	"net/http"
	"github.com/gorilla/mux"
	merchantPb "github.com/rumsrami/givly-rpc-api/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/rpc/transaction"
)

// Server represents a twirp RPC server
type Server struct {
	Router *mux.Router
}

// NewRPCServer returns a new Twirp RPC server
func NewRPCServer() *Server{
	r := mux.NewRouter()
	route(r)
	server := &Server{
		Router: r,
	}
	return server
}

// route mounts the Services on the Server router
func route(r *mux.Router) {
	routeMerchant(r)
	routeTransaction(r)
}

// Run runs the RPC server
func (s *Server) Run(addr string) {
	http.ListenAndServe(addr, s.Router)
}


// routeMerchant mounts Merchant server Handler
func routeMerchant(r *mux.Router) {
	mdServer := NewMerchantDirectory()
	twirpmdHandler := merchantPb.NewMerchantServiceServer(mdServer, nil)
	r.PathPrefix(merchantPb.MerchantServicePathPrefix).Handler(twirpmdHandler)
}

// routeTransaction mounts Transaction server Handler
func routeTransaction(r *mux.Router) {
	tdServer := NewTransactionDirectory()
	twirptdHandler := transactionPb.NewTransactionServiceServer(tdServer, nil)
	r.PathPrefix(transactionPb.TransactionServicePathPrefix).Handler(twirptdHandler)
}
