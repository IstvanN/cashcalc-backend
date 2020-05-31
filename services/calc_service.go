/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"fmt"
	"math"

	"github.com/IstvanN/cashcalc-backend/models"
)

// IsZoneEU returns if zone is EU or not
func isZoneEU(zn int) bool {
	return zn <= 4 && zn >= 0
}

// CalcBaseFareWithVatAndDiscountAir calculates the basefare increased by VAT and applied discount
func CalcBaseFareWithVatAndDiscountAir(inputData models.CalcInputData, vatPercent float64, baseFare int) float64 {
	if isZoneEU(inputData.ZoneNumber) {
		baseFareIncreasedWithVat := IncreaseWithVat(float64(baseFare), vatPercent)
		return math.Round(applyDiscountToBaseFare(baseFareIncreasedWithVat, inputData.DiscountPercent))
	}

	if inputData.IsDocument && inputData.Weight <= 2 {
		return math.Round(applyDiscountToBaseFare(float64(baseFare), inputData.DiscountPercent))
	}

	return math.Round(applyDiscountToBaseFare(float64(baseFare), inputData.DiscountPercent))
}

// TODO: WRITE TEST

// ValidateInputData takes an input data model and returns with an error if there is a logical error
func ValidateInputData(input models.CalcInputData) error {
	var err error
	if isZoneEU(input.ZoneNumber) && input.IsDocument {
		err = fmt.Errorf("zone number %v, document status %v: no document delivery to EU", input.ZoneNumber, input.IsDocument)
		return err
	}

	if input.IsDocument && input.Weight > 2 {
		err = fmt.Errorf("weight %v, document status %v: document cannot have more weight than 2", input.Weight, input.IsDocument)
	}
	return nil
}

func applyDiscountToBaseFare(baseFare float64, discountPercent float64) float64 {
	return math.Round((1 - discountPercent/100) * baseFare)
}
