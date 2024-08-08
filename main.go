package main

import (
	"fmt"

	"example.com/price_calculator/cmdmanager"
	"example.com/price_calculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncluidedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
