package http

import (
	"net/http"
	"github.com/gorilla/mux"
	merchantPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/merchant"
	transactionPb "github.com/rumsrami/givly-rpc-api/pkg/rpc/transaction"
	"github.com/rumsrami/givly-rpc-api/pkg/adding"
	"github.com/rumsrami/givly-rpc-api/pkg/listing"
)

// Server represents a twirp RPC server
type Server struct {
	Router *mux.Router
}

// NewRPCServer returns a new Twirp RPC server
func NewRPCServer(a adding.Service, l listing.Service) *Server {
	r := mux.NewRouter()
	route(r, a, l)
	server := &Server{
		Router: r,
	}
	return server
}

// route mounts the Services on the Server router
// route adds dependencies on routes
func route(r *mux.Router, a adding.Service, l listing.Service) {
	routeMerchant(r, a, l)
	routeTransaction(r, a, l)
}

// Run runs the RPC server
func (s *Server) Run(addr string) {
	http.ListenAndServe(addr, s.Router)
}


// routeMerchant mounts Merchant server Handler
func routeMerchant(r *mux.Router, a adding.Service, l listing.Service) {
	mdServer := NewMerchantDirectory(a, l)
	twirpmdHandler := merchantPb.NewMerchantServiceServer(mdServer, nil)
	r.PathPrefix(merchantPb.MerchantServicePathPrefix).Handler(twirpmdHandler)
}

// routeTransaction mounts Transaction server Handler
func routeTransaction(r *mux.Router, a adding.Service, l listing.Service) {
	tdServer := NewTransactionDirectory(a, l)
	twirptdHandler := transactionPb.NewTransactionServiceServer(tdServer, nil)
	r.PathPrefix(transactionPb.TransactionServicePathPrefix).Handler(twirptdHandler)
}
