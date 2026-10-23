package lsp

import "fmt"

type PricingStrategy interface {
	CalculateFee(hours int) (float64, error)
}

type CarPricing struct{}

func (cp *CarPricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid hours %d", hours)
	}
	return float64(hours) * 10, nil
}

type MotorcyclePricing struct{}

func (mp *MotorcyclePricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid hours %d", hours)
	}
	return float64(hours) * 5, nil
}

type TruckPricing struct{}

func (tp *TruckPricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid hours %d", hours)
	}
	return float64(hours) * 20, nil
}
