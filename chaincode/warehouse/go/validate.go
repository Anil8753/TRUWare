package main

import (
	"errors"
	"fmt"
)

func ValidateAssetData(asset *Asset) error {

	if asset.Id == "" ||
		asset.Postion.Latitude == 0 ||
		asset.Postion.Longitude == 0 ||
		asset.General.Address == "" ||
		asset.General.Name == "" ||
		asset.General.TotalArea == 0 ||
		asset.General.Details == "" {

		return errors.New(`these are mandatory fields.
						asset.Id, 
						asset.Postion.Latitude,
						asset.Postion.Longitude,
						asset.General.Address,
						asset.General.Name,
						asset.General.TotalArea,
						asset.General.Details`)
	}

	if asset.General.TotalArea < asset.General.AllocatedArea {
		return fmt.Errorf(
			"occupied area cannot be more than total area. occupied area: %d, total area: %d",
			asset.General.TotalArea, asset.General.AllocatedArea,
		)
	}

	if asset.General.Rate < 1 {
		return fmt.Errorf("rate cannot be less than 1. Rate: %d", asset.General.Rate)
	}

	return nil
}
