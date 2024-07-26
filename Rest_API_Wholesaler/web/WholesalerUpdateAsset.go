package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (setup *OrgSetup) WholesalerUpdateAsset(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received WholesalerUpdateAsset request")

	// Define a structure for the expected JSON payload
	type Request struct {
		ID                string `json:"id"`
		WholesalerId      string `json:"wholesalerId"`
		WholesalerName    string `json:"wholesalerName"`
		WholesalerBuyDate string `json:"wholesalerBuyDate"`
	}

	var requestData Request
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "JSON Decode error: "+err.Error(), http.StatusBadRequest)
		return
	}

	network := setup.Gateway.GetNetwork(setup.Channel)
	contract := network.GetContract(setup.Chaincode)

	// Retrieve existing asset data
	result, err := contract.EvaluateTransaction("ReadAsset", requestData.ID)
	if err != nil {
		http.Error(w, "Error invoking ReadAsset: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshal the asset data to update it
	var asset map[string]string
	if err := json.Unmarshal(result, &asset); err != nil {
		http.Error(w, "JSON Unmarshal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the asset with the new wholesaler information
	if requestData.WholesalerId != "" {
		asset["WholesalerId"] = requestData.WholesalerId
	}
	if requestData.WholesalerName != "" {
		asset["WholesalerName"] = requestData.WholesalerName
	}
	if requestData.WholesalerBuyDate != "" {
		asset["WholesalerBuyDate"] = requestData.WholesalerBuyDate
	}
	// Submit transaction to the ledger to update the asset
	_, err = contract.SubmitTransaction("UpdateAsset", requestData.ID, asset["FarmerId"], asset["FarmLocation"], asset["Variety"], asset["BatchNo"], asset["HarvestDate"], asset["Price"], asset["Quantity"], asset["WholesalerId"], asset["WholesalerPrice"], asset["WholesalerBuyDate"], asset["RetailerId"], asset["RetailerBuyDate"])
	if err != nil {
		http.Error(w, "Error invoking UpdateAsset: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response with the updated asset ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Asset updated successfully", "id": requestData.ID})
}
