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
	FarmLocation      string `json:"FarmLocation"`
	Variety           string `json:"Variety"`
	BatchNo           string `json:"BatchNo"`
	HarvestDate       string `json:"HarvestDate"`
	Price             string `json:"Price"`
	Quantity          string `json:"Quantity"`
	WholesalerId      string `json:"WholesalerId"`
	WholesalerPrice   string `json:"WholesalerPrice"`
	WholesalerBuyDate string `json:"WholesalerBuyDate"`
	RetailerId        string `json:"RetailerId"`
	RetailerBuyDate   string `json:"RetailerBuyDate"`
}

// InitLedger initializes the ledger with a set of sample assets
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "asset1", FarmerId: "F001", FarmLocation: "Location1", Variety: "Variety1", BatchNo: "B001", HarvestDate: "2023-01-01", Price: "1000", Quantity: "100"},
		{ID: "asset2", FarmerId: "F002", FarmLocation: "Location2", Variety: "Variety2", BatchNo: "B002", HarvestDate: "2023-02-01", Price: "1100", Quantity: "150"},
		{ID: "asset3", FarmerId: "F003", FarmLocation: "Location3", Variety: "Variety3", BatchNo: "B003", HarvestDate: "2023-03-01", Price: "1200", Quantity: "200"},
		{ID: "asset4", FarmerId: "F004", FarmLocation: "Location4", Variety: "Variety4", BatchNo: "B004", HarvestDate: "2023-04-01", Price: "1300", Quantity: "250"},
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
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id, farmerId, farmLocation, variety, batchNo, harvestDate, price, quantity, wholesalerId, wholesalerPrice, wholesalerBuyDate, retailerId, retailerBuyDate string) error {
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
		FarmLocation:      farmLocation,
		Variety:           variety,
		BatchNo:           batchNo,
		HarvestDate:       harvestDate,
		Price:             price,
		Quantity:          quantity,
		WholesalerId:      wholesalerId,
		WholesalerPrice:   wholesalerPrice,
		WholesalerBuyDate: wholesalerBuyDate,
		RetailerId:        retailerId,
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
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id, farmerId, farmLocation, variety, batchNo, harvestDate, price, quantity, wholesalerId, wholesalerPrice, wholesalerBuyDate, retailerId, retailerBuyDate string) error {
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
		FarmLocation:      farmLocation,
		Variety:           variety,
		BatchNo:           batchNo,
		HarvestDate:       harvestDate,
		Price:             price,
		Quantity:          quantity,
		WholesalerId:      wholesalerId,
		WholesalerPrice:   wholesalerPrice,
		WholesalerBuyDate: wholesalerBuyDate,
		RetailerId:        retailerId,
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
