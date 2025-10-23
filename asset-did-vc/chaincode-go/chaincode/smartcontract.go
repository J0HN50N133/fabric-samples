package chaincode

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

const accumulatorKey = "accumulator"

// SmartContract provides functions for managing an Accumulator
type SmartContract struct {
	contractapi.Contract
}

// Accumulator stores the state of the cryptographic accumulator on the ledger.
// Digest is stored as raw bytes.
type Accumulator struct {
	Digest  []byte `json:"digest"`
	Version uint64 `json:"version"`
}

// AccumulatorResponse is the structure returned to the client.
// Digest is hex-encoded for client convenience.
type AccumulatorResponse struct {
	Digest  string `json:"digest"`
	Version uint64 `json:"version"`
}

// InitLedger creates the initial accumulator on the ledger.
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	exists, err := s.accumulatorExists(ctx)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("accumulator has already been initialized")
	}

	// Initial digest is 32 bytes of zeros, version is 1.
	digest := make([]byte, 32)
	acc := Accumulator{
		Digest:  digest,
		Version: 1,
	}
	accJSON, err := json.Marshal(acc)
	if err != nil {
		return fmt.Errorf("failed to marshal initial accumulator: %v", err)
	}

	err = ctx.GetStub().PutState(accumulatorKey, accJSON)
	if err != nil {
		return fmt.Errorf("failed to put initial accumulator to world state: %v", err)
	}

	return nil
}

// UpdateAccumulator updates the accumulator's digest and increments the version.
func (s *SmartContract) UpdateAccumulator(ctx contractapi.TransactionContextInterface, newDigestHex string) error {
	acc, err := s.getAccumulatorState(ctx)
	if err != nil {
		return err
	}

	newDigest, err := hex.DecodeString(newDigestHex)
	if err != nil {
		return fmt.Errorf("failed to decode hex string for new digest: %v", err)
	}
	if len(newDigest) != 32 {
		return fmt.Errorf("new digest must be 32 bytes, but got %d bytes", len(newDigest))
	}

	acc.Digest = newDigest
	acc.Version++

	accJSON, err := json.Marshal(acc)
	if err != nil {
		return fmt.Errorf("failed to marshal updated accumulator: %v", err)
	}

	return ctx.GetStub().PutState(accumulatorKey, accJSON)
}

// GetAccumulator returns the current accumulator's digest (hex-encoded) and version.
func (s *SmartContract) GetAccumulator(ctx contractapi.TransactionContextInterface) (*AccumulatorResponse, error) {
	acc, err := s.getAccumulatorState(ctx)
	if err != nil {
		return nil, err
	}

	response := &AccumulatorResponse{
		Digest:  hex.EncodeToString(acc.Digest),
		Version: acc.Version,
	}

	return response, nil
}

// getAccumulatorState retrieves the accumulator from the world state.
func (s *SmartContract) getAccumulatorState(ctx contractapi.TransactionContextInterface) (*Accumulator, error) {
	accJSON, err := ctx.GetStub().GetState(accumulatorKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read accumulator from world state: %v", err)
	}
	if accJSON == nil {
		return nil, fmt.Errorf("accumulator has not been initialized")
	}

	var acc Accumulator
	err = json.Unmarshal(accJSON, &acc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal accumulator data: %v", err)
	}

	return &acc, nil
}

// accumulatorExists checks if the accumulator has been initialized in the world state.
func (s *SmartContract) accumulatorExists(ctx contractapi.TransactionContextInterface) (bool, error) {
	accJSON, err := ctx.GetStub().GetState(accumulatorKey)
	if err != nil {
		return false, fmt.Errorf("failed to read accumulator from world state: %v", err)
	}
	return accJSON != nil, nil
}
