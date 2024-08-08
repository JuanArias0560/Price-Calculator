package prices

import (
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/iomanager"
)

type TaxIncluidedPriceJob struct {
	IOManager          iomanager.IOManager `json:"-"`
	TaxRate            float64             `json:"tax_rate"`
	InputPrices        []float64           `json:"input_prices"`
	TaxIncluidedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncluidedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncluidedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncluidedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncluidedPrice)
	}

	job.TaxIncluidedPrices = result

	return job.IOManager.WriteResult(job)

}

func NewTaxIncluidedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncluidedPriceJob {
	return &TaxIncluidedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
