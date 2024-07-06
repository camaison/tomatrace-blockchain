package web

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
	Chaincode    string
	Channel      string
}

var clientsMutex sync.Mutex

// Serve initializes and starts the HTTP server
func Serve(setups OrgSetup) {
	mux := http.NewServeMux()

	// Define routes for direct endpoints
	mux.HandleFunc("/retailerUpdate", setups.RetailerUpdateAsset)
	mux.HandleFunc("/getAll", setups.GetAllAssets)
	mux.HandleFunc("/getEntry", setups.ReadAsset)

	// Wrap the mux with the logging middleware
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Listening on http://localhost:3002/ ...")
	if err := http.ListenAndServe(":3002", loggedMux); err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}

// loggingMiddleware logs all incoming HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
