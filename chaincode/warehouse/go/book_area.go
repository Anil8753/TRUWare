package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// BookArea books the specified area from the total available space.
func (s *WarehouseContract) BookArea(
	ctx contractapi.TransactionContextInterface,
	id string,
	bookingArea int,
	duration int,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	available := asset.TotalArea - asset.BookedArea

	if (available) < bookingArea {
		return fmt.Errorf(
			"not enough space. available area: %d,  requested booking area %d",
			available, bookingArea,
		)
	}

	asset.Bookings = append(asset.Bookings, Booking{
		BookerID: "xyz",
		Area:     bookingArea,
		Duration: duration,
	})

	asset.BookedArea += bookingArea

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
