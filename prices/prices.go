package prices

import "fmt"

type TaxIncluidedPriceJob struct {
	TaxRate            float64
	InputPrices        []float64
	TaxIncluidedPrices map[string]float64
}

func (job TaxIncluidedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintln("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func NewTaxIncluidedPriceJob(taxRate float64) *TaxIncluidedPriceJob {
	return &TaxIncluidedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
