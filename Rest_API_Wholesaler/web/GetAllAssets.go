package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (setup *OrgSetup) GetAllAssets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Get All Assets request")

	network := setup.Gateway.GetNetwork(setup.Channel)
	contract := network.GetContract(setup.Chaincode)

	// Evaluate transaction using the GetAllData function from chaincode
	result, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		fmt.Fprintf(w, "Error querying GetAllAssets: %s", err)
		return
	}

	// Prepare the response to return JSON data
	var data interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		fmt.Fprintf(w, "Error unmarshaling JSON data: %s", err)
		return
	}

	// Send the response with data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
