package main

import (
	"fmt"

	"example.com/price-calulator/prices"
)

func main() {
	
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob:= prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(result)
}