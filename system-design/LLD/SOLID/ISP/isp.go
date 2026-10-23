/*
Refactor the following interface:
type Vehicle interface {
    CanFit(spot *ParkingSpot) bool
    CalculateFee(hours int) (float64, error)
    ProcessPayment(amount float64) error
}

Create the three smaller interfaces and ensure that:
- Car and Motorcycle implement Vehicle.
- CarPricing, MotorcyclePricing, and TruckPricing implement FeeCalculator.
- A PaymentService implements PaymentProcessor.
*/

package isp

import (
	"fmt"
	parkinglot "solid/lot"
)

type Vehicle interface {
	CanFit(spot *parkinglot.ParkingSpot) bool
}

type Car struct{}

func (c *Car) CanFit(spot *parkinglot.ParkingSpot) bool {
	if spot.ID == "A101" {
		return true
	}
	return false
}

type Motorcycle struct{}

func (m *Motorcycle) CanFit(spot *parkinglot.ParkingSpot) bool {
	if spot.ID == "B101" {
		return true
	}
	return false
}

type FeeCalculator interface {
	CalculateFee(hours int) (float64, error)
}

type CarPricing struct{}

func (cp *CarPricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid Hours %d", hours)
	}
	return float64(hours) * 10, nil
}

type MotorcyclePricing struct{}

func (mp *MotorcyclePricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid Hours %d", hours)
	}
	return float64(hours) * 5, nil
}

type TruckPricing struct{}

func (tp *TruckPricing) CalculateFee(hours int) (float64, error) {
	if hours < 0 {
		return 0, fmt.Errorf("Invalid Hours %d", hours)
	}
	return float64(hours) * 20, nil
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	fmt.Println("Processed amount", amount, "successfully")
	return nil
}

/*
Interface consumers
Let's verify that the interfaces are useful to callers, not just correctly declared.

Write these two functions:
func GetFee(calc FeeCalculator, hours int) (float64, error) {
    // Calculate the fee through the interface.
}

func MakePayment(processor PaymentProcessor, amount float64) error {
    // Process payment through the interface.
}

Constraints:
- Don't accept concrete types like *CarPricing or *PaymentService.
- Don't use type assertions or switch.
- Test GetFee with car and truck pricing.
- Test MakePayment with PaymentService.
*/

func GetFee(calc FeeCalculator, hours int) (float64, error) {
	return calc.CalculateFee(hours)
}

func MakePayment(processor PaymentProcessor, amount float64) error {
	return processor.ProcessPayment(amount)
}

func ISP() {
	tp := TruckPricing{}

	fee, err := GetFee(&tp, 2)
	if err != nil {
		fmt.Println("Error getting fee", err)
		return
	}
	fmt.Println("Truck Fee", fee)

	ps := PaymentService{}

	err = MakePayment(&ps, fee)
	if err != nil {
		fmt.Println("Error making payment", err)
		return
	}
}
