package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing assets
type SmartContract struct {
	contractapi.Contract
}

// Asset represents the structure for an asset on the ledger
type Asset struct {
	ID                string `json:"ID"`
	FarmerId          string `json:"FarmerId"`
	FarmerName        string `json:"FarmerName"`
	FarmLocation      string `json:"FarmLocation"`
	Variety           string `json:"Variety"`
	BatchNo           string `json:"BatchNo"`
	HarvestDate       string `json:"HarvestDate"`
	Price             string `json:"Price"`
	Quantity          string `json:"Quantity"`
	WholesalerId      string `json:"WholesalerId"`
	WholesalerName    string `json:"WholesalerName"`
	WholesalerBuyDate string `json:"WholesalerBuyDate"`
	RetailerId        string `json:"RetailerId"`
	RetailerName      string `json:"RetailerName"`
	RetailerBuyDate   string `json:"RetailerBuyDate"`
}

// InitLedger initializes the ledger with a set of sample assets
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "1", FarmerId: "1", FarmerName: "Farmer 1", FarmLocation: "Location 1", Variety: "Variety 1", BatchNo: "Batch 1", HarvestDate: "2021-01-01", Price: "100", Quantity: "100", WholesalerId: "2", WholesalerName: "Wholesaler 1", WholesalerBuyDate: "2021-01-02", RetailerId: "3", RetailerName: "Retailer 1", RetailerBuyDate: "2021-01-03"},
		{ID: "2", FarmerId: "2", FarmerName: "Farmer 2", FarmLocation: "Location 2", Variety: "Variety 2", BatchNo: "Batch 2", HarvestDate: "2021-02-01", Price: "200", Quantity: "200", WholesalerId: "3", WholesalerName: "Wholesaler 2", WholesalerBuyDate: "2021-02-02", RetailerId: "4", RetailerName: "Retailer 2", RetailerBuyDate: "2021-02-03"},
		{ID: "3", FarmerId: "3", FarmerName: "Farmer 3", FarmLocation: "Location 3", Variety: "Variety 3", BatchNo: "Batch 3", HarvestDate: "2021-03-01", Price: "300", Quantity: "300", WholesalerId: "4", WholesalerName: "Wholesaler 3", WholesalerBuyDate: "2021-03-02", RetailerId: "5", RetailerName: "Retailer 3", RetailerBuyDate: "2021-03-03"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put asset %s: %v", asset.ID, err)
		}
	}

	return nil
}

// CreateAsset creates a new asset and stores it in the ledger
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id, farmerId, farmerName, farmLocation, variety, batchNo, harvestDate, price, quantity, wholesalerId, WholesalerName, wholesalerPrice, wholesalerBuyDate, retailerId, retailerName, retailerBuyDate string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID:                id,
		FarmerId:          farmerId,
		FarmerName:        farmerName,
		FarmLocation:      farmLocation,
		Variety:           variety,
		BatchNo:           batchNo,
		HarvestDate:       harvestDate,
		Price:             price,
		Quantity:          quantity,
		WholesalerId:      wholesalerId,
		WholesalerName:    WholesalerName,
		WholesalerBuyDate: wholesalerBuyDate,
		RetailerId:        retailerId,
		RetailerName:      retailerName,
		RetailerBuyDate:   retailerBuyDate,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset retrieves an asset from the ledger by its ID
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read asset %s from world state: %v", id, err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the ledger
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id, farmerId, farmerName, farmLocation, variety, batchNo, harvestDate, price, quantity, wholesalerId, WholesalerName, wholesalerPrice, wholesalerBuyDate, retailerId, retailerName, retailerBuyDate string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset := Asset{
		ID:                id,
		FarmerId:          farmerId,
		FarmerName:        farmerName,
		FarmLocation:      farmLocation,
		Variety:           variety,
		BatchNo:           batchNo,
		HarvestDate:       harvestDate,
		Price:             price,
		Quantity:          quantity,
		WholesalerId:      wholesalerId,
		WholesalerName:    WholesalerName,
		WholesalerBuyDate: wholesalerBuyDate,
		RetailerId:        retailerId,
		RetailerName:      retailerName,
		RetailerBuyDate:   retailerBuyDate,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// AssetExists checks if an asset exists in the ledger
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, err
	}

	return assetJSON != nil, nil
}

// GetAllAssets retrieves all assets from the ledger
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
