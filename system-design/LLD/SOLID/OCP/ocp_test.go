/*
Write a test function that checks all three pricing implementations for these inputs:

Strategy	Hours	Expected fee
Car			  0			0
Car			  2			20
Motorcycle	  0			0
Motorcycle	  2			10
Truck		  0			0
Truck		  2			40

Syntax you'll need
1. Define a test
func TestPricingContract(t *testing.T) {
    // assertions
}

2. Compare actual and expected values
if got != want {
    t.Errorf("got %v, want %v", got, want)
}

Constraints
- Use the PricingStrategy interface so every implementation is tested through the same contract.
- Don't use a switch statement.
- Don't modify the pricing implementations just to make the test pass.
- Use a table-driven test or another approach that avoids unnecessary duplication.
*/

package ocp

import "testing"

type TestCase struct {
	Name     string
	Strategy PricingStrategy
	Hours    int
	want     float64
}

func TestPricingContract(t *testing.T) {
	testCases := []TestCase{
		{
			Name:     "car-0",
			Strategy: &CarPricing{},
			Hours:    0,
			want:     0,
		},
		{
			Name:     "car-2h",
			Strategy: &CarPricing{},
			Hours:    2,
			want:     20,
		},
		{
			Name:     "bike-0",
			Strategy: &MotorcyclePricing{},
			Hours:    0,
			want:     0,
		},
		{
			Name:     "bike-2h",
			Strategy: &MotorcyclePricing{},
			Hours:    2,
			want:     10,
		},
		{
			Name:     "truck-0",
			Strategy: &TruckPricing{},
			Hours:    0,
			want:     0,
		},
		{
			Name:     "truck-2h",
			Strategy: &TruckPricing{},
			Hours:    2,
			want:     40,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Strategy.CalculateFee(tc.Hours)

			if got != tc.want {
				t.Errorf("expected %.2f, got %.2f", tc.want, got)
			}
		})
	}
}
