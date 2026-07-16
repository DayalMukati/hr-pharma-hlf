package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for pharma cold-chain custody tracking
type SmartContract struct {
	contractapi.Contract
}

// Shipment represents a temperature-sensitive pharmaceutical shipment.
type Shipment struct {
	ShipmentID string `json:"ShipmentID"` // unique shipment id, e.g. "shp1"
	Drug       string `json:"Drug"`       // drug name
	Custodian  string `json:"Custodian"`  // party currently holding the shipment
	Breached   bool   `json:"Breached"`   // true once a cold-chain breach is recorded
	Status     string `json:"Status"`     // IN_TRANSIT | DELIVERED
}

// HistoryEntry represents one revision of a shipment from the ledger history.
type HistoryEntry struct {
	TxID      string    `json:"TxID"`
	Value     *Shipment `json:"Value"`
	Timestamp string    `json:"Timestamp"`
	IsDelete  bool      `json:"IsDelete"`
}

// CreateShipment registers a new shipment held by the origin custodian, with
// status "IN_TRANSIT" and no breach.
// It must fail if the shipment already exists.
func (s *SmartContract) CreateShipment(ctx contractapi.TransactionContextInterface, shipmentID string, drug string, origin string) error {

	return nil
}

// GetShipment returns the shipment identified by shipmentID.
// It must fail if the shipment does not exist.
func (s *SmartContract) GetShipment(ctx contractapi.TransactionContextInterface, shipmentID string) (*Shipment, error) {

	return nil, nil
}

// TransferCustody hands the shipment to a new custodian.
// It must fail if the shipment does not exist, is already DELIVERED, or
// newCustodian is empty.
func (s *SmartContract) TransferCustody(ctx contractapi.TransactionContextInterface, shipmentID string, newCustodian string) error {

	return nil
}

// RecordBreach flags a cold-chain temperature breach on the shipment. Once set,
// the breach flag can never be cleared.
// It must fail if the shipment does not exist.
func (s *SmartContract) RecordBreach(ctx contractapi.TransactionContextInterface, shipmentID string) error {

	return nil
}

// DeliverShipment marks the shipment as "DELIVERED".
// It must fail if the shipment does not exist or is already DELIVERED.
func (s *SmartContract) DeliverShipment(ctx contractapi.TransactionContextInterface, shipmentID string) error {

	return nil
}

// GetProvenance returns the full custody history of a shipment, newest first,
// using GetHistoryForKey.
func (s *SmartContract) GetProvenance(ctx contractapi.TransactionContextInterface, shipmentID string) ([]HistoryEntry, error) {

	return nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic("Error creating pharma chaincode: " + err.Error())
	}

	if err := chaincode.Start(); err != nil {
		panic("Error starting pharma chaincode: " + err.Error())
	}
}
