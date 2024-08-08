package prices

import (
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/filemanager"
)

type TaxIncluidedPriceJob struct {
	IOManager          filemanager.Filemanager `json:"-"`
	TaxRate            float64                 `json:"tax_rate"`
	InputPrices        []float64               `json:"input_prices"`
	TaxIncluidedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncluidedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncluidedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncluidedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncluidedPrice)
	}

	job.TaxIncluidedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIncluidedPriceJob(fm filemanager.Filemanager, taxRate float64) *TaxIncluidedPriceJob {
	return &TaxIncluidedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
