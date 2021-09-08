package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// FreeArea releases the specified space.
func (s *WarehouseContract) FreeArea(
	ctx contractapi.TransactionContextInterface,
	id string,
	bookingID string,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	index := -1
	bookingArea := 0

	for _, booking := range asset.Bookings {
		if booking.ID == bookingID {
			bookingArea = booking.Area
			break
		}
		index++
	}

	if index != -1 {
		return fmt.Errorf("booking ID not found: %s", bookingID)
	}

	asset.BookedArea -= bookingArea
	asset.Bookings = append(asset.Bookings[:index], asset.Bookings[index+1:]...)

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
