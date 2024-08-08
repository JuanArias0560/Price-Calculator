package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price_calculator/conversion"
)

type TaxIncluidedPriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TaxIncluidedPrices map[string]float64
}

func (job *TaxIncluidedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could nor open file!")
		fmt.Println(err)
		return
	}

	scaner := bufio.NewScanner(file)

	var lines []string

	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}

	err = scaner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()

}

func (job *TaxIncluidedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncluidedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncluidedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncluidedPriceJob(taxRate float64) *TaxIncluidedPriceJob {
	return &TaxIncluidedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
