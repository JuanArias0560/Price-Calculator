package main

import (
	"fmt"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncluidedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
