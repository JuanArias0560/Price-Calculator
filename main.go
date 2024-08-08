package main

import (
	"fmt"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncluidedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-errorChans[index]:
			fmt.Println("Done!")

		}
	}
}
