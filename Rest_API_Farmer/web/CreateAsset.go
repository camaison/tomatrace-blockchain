package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (setup *OrgSetup) CreateAsset(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received CreateAsset request")

	// Define a structure for the expected JSON payload
	type Request struct {
		ID           string `json:"id"`
		FarmerId     string `json:"farmerId"`
		FarmerName   string `json:"farmerName"`
		FarmLocation string `json:"farmLocation"`
		Variety      string `json:"variety"`
		BatchNo      string `json:"batchNo"`
		HarvestDate  string `json:"harvestDate"`
		Price        string `json:"price"`
		Quantity     string `json:"quantity"`
	}

	var requestData Request
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "JSON Decode error: "+err.Error(), http.StatusBadRequest)
		return
	}

	network := setup.Gateway.GetNetwork(setup.Channel)
	contract := network.GetContract(setup.Chaincode)

	// Submit transaction to the ledger to create the asset
	_, err := contract.SubmitTransaction("CreateAsset", requestData.ID, requestData.FarmerId, requestData.FarmLocation, requestData.Variety, requestData.BatchNo, requestData.HarvestDate, requestData.Price, requestData.Quantity, "", "", "", "", "")
	if err != nil {
		http.Error(w, "Error invoking CreateAsset: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response with the asset ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Asset created successfully", "id": requestData.ID})
}
