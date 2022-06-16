package main

import "fmt"

// Chapter six

func testFunctions() {
	dolarPrice, brlRealPrice := FinalPrice(34.99)

	fmt.Printf("Final price in Dolar: %.2f\nFinal price in Real: %.2f\n", dolarPrice, brlRealPrice)
}

func FinalPrice(priceCost float64) (dolarPrice float64, brlRealPrice float64) {
	profitFactor := 1.33
	conversionTax := 5.05

	dolarPrice = priceCost * profitFactor
	brlRealPrice = dolarPrice * conversionTax

	return dolarPrice, brlRealPrice
}
