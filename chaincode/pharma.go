package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for pharma cold-chain custody tracking
type SmartContract struct {
	contractapi.Contract
}

// Shipment represents a temperature-sensitive pharmaceutical shipment.
type Shipment struct {
	ShipmentID string `json:"ShipmentID"`
	Drug       string `json:"Drug"`
	Custodian  string `json:"Custodian"`
	Breached   bool   `json:"Breached"`
	Status     string `json:"Status"`
}

// HistoryEntry represents one revision of a shipment from the ledger history.
type HistoryEntry struct {
	TxID      string    `json:"TxID"`
	Value     *Shipment `json:"Value"`
	Timestamp string    `json:"Timestamp"`
	IsDelete  bool      `json:"IsDelete"`
}

const (
	statusInTransit = "IN_TRANSIT"
	statusDelivered = "DELIVERED"
)

// CreateShipment registers a new shipment held by the origin custodian.
func (s *SmartContract) CreateShipment(ctx contractapi.TransactionContextInterface, shipmentID string, drug string, origin string) error {
	existing, err := ctx.GetStub().GetState(shipmentID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if existing != nil {
		return fmt.Errorf("shipment %s already exists", shipmentID)
	}

	shipment := Shipment{
		ShipmentID: shipmentID,
		Drug:       drug,
		Custodian:  origin,
		Breached:   false,
		Status:     statusInTransit,
	}
	return putShipment(ctx, &shipment)
}

// GetShipment returns the shipment identified by shipmentID.
func (s *SmartContract) GetShipment(ctx contractapi.TransactionContextInterface, shipmentID string) (*Shipment, error) {
	data, err := ctx.GetStub().GetState(shipmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if data == nil {
		return nil, fmt.Errorf("shipment %s does not exist", shipmentID)
	}

	var shipment Shipment
	if err := json.Unmarshal(data, &shipment); err != nil {
		return nil, err
	}
	return &shipment, nil
}

// TransferCustody hands the shipment to a new custodian.
func (s *SmartContract) TransferCustody(ctx contractapi.TransactionContextInterface, shipmentID string, newCustodian string) error {
	if newCustodian == "" {
		return fmt.Errorf("new custodian cannot be empty")
	}

	shipment, err := s.GetShipment(ctx, shipmentID)
	if err != nil {
		return err
	}
	if shipment.Status == statusDelivered {
		return fmt.Errorf("shipment %s is already DELIVERED and cannot change custody", shipmentID)
	}

	shipment.Custodian = newCustodian
	return putShipment(ctx, shipment)
}

// RecordBreach flags a cold-chain breach. The flag is one-way.
func (s *SmartContract) RecordBreach(ctx contractapi.TransactionContextInterface, shipmentID string) error {
	shipment, err := s.GetShipment(ctx, shipmentID)
	if err != nil {
		return err
	}

	shipment.Breached = true
	return putShipment(ctx, shipment)
}

// DeliverShipment marks the shipment as "DELIVERED".
func (s *SmartContract) DeliverShipment(ctx contractapi.TransactionContextInterface, shipmentID string) error {
	shipment, err := s.GetShipment(ctx, shipmentID)
	if err != nil {
		return err
	}
	if shipment.Status == statusDelivered {
		return fmt.Errorf("shipment %s is already DELIVERED", shipmentID)
	}

	shipment.Status = statusDelivered
	return putShipment(ctx, shipment)
}

// GetProvenance returns the full custody history of a shipment, newest first.
func (s *SmartContract) GetProvenance(ctx contractapi.TransactionContextInterface, shipmentID string) ([]HistoryEntry, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(shipmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for %s: %v", shipmentID, err)
	}
	defer resultsIterator.Close()

	var history []HistoryEntry
	for resultsIterator.HasNext() {
		modification, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		entry := HistoryEntry{
			TxID:      modification.TxId,
			Timestamp: time.Unix(modification.Timestamp.Seconds, int64(modification.Timestamp.Nanos)).UTC().Format(time.RFC3339),
			IsDelete:  modification.IsDelete,
		}
		if !modification.IsDelete {
			var shipment Shipment
			if err := json.Unmarshal(modification.Value, &shipment); err != nil {
				return nil, err
			}
			entry.Value = &shipment
		}
		history = append(history, entry)
	}
	return history, nil
}

// --- helpers ---

func putShipment(ctx contractapi.TransactionContextInterface, shipment *Shipment) error {
	bytes, err := json.Marshal(shipment)
	if err != nil {
		return err
	}
	return ctx.GetStub().PutState(shipment.ShipmentID, bytes)
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
