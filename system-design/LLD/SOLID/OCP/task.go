/*
add truck pricing
Implement a TruckPricing type that satisfies PricingStrategy and charges ₹20 per hour.

Constraints:
- Don't modify PricingStrategy.
- Don't modify CarPricing or MotorcyclePricing.
- Don't add a switch statement.
- Use the same method signature as the interface.
*/

package ocp

import "fmt"

type PricingStrategy interface {
	CalculateFee(hours int) float64
}

type CarPricing struct{}

func (cp *CarPricing) CalculateFee(hours int) float64 {
	return float64(hours) * 10
}

type MotorcyclePricing struct{}

func (mp *MotorcyclePricing) CalculateFee(hours int) float64 {
	return float64(hours) * 5
}

type TruckPricing struct{}

func (tp *TruckPricing) CalculateFee(hours int) float64 {
	return float64(hours) * 20
}

/*
Write this function:
func PrintFee(pricing PricingStrategy, hours int) {
    //Calculate and print the fee using the interface.
}

Requirements
1. Accept a PricingStrategy, not a concrete pricing type.
2. Calculate the fee through the interface.
3. Test it with car, motorcycle, and truck pricing for 2 hours.
4. Don't use switch or type assertions.
*/

func PrintFee(ps PricingStrategy, hours int) {
	fmt.Println(ps.CalculateFee(hours))
}

func OCP() {
	cp := CarPricing{}
	mp := MotorcyclePricing{}
	tp := TruckPricing{}

	PrintFee(&cp, 2)
	PrintFee(&mp, 2)
	PrintFee(&tp, 2)
}
