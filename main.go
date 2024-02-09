package main

import "fmt"

func main() {
	var prices []float64 = []float64{10, 20, 30}
	var taxRates []float64 = []float64{0, 0.07, 0.1, 0.15}

	var result map[float64][]float64 = make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for priceindex, price := range prices {
			taxIncludedPrices[priceindex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}