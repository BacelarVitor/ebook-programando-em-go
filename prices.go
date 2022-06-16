package main 

import fmt 

// Chapter six

func finalPrice(priceCost float64) (dolarPrice float64, brlRealPrice float64) {
	profitFactor := 1.33
	conversionTax := 5.05

	dolarPrice = priceCost * profitFactor
	brlRealPrice = dolarPrice * conversionTax

	return dolarPrice, brlRealPrice
}