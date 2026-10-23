/*
Refactor ParkingCheckout so it depends on your existing PaymentProcessor interface.

type ParkingCheckout struct{}

func (c *ParkingCheckout) Checkout(amount float64) error {
    processor := PaymentService{}
    return processor.ProcessPayment(amount)
}

Your task: Modify ParkingCheckout so that:
- It receives a PaymentProcessor from outside instead of creating PaymentService internally.
- Checkout uses the injected processor.
- You don't need to change Checkout when switching payment implementations.
*/

package dip

import (
	"errors"
	"fmt"
	isp "solid/ISP"
)

type ParkingCheckout struct {
	Processor isp.PaymentProcessor
}

func (c *ParkingCheckout) Checkout(amount float64) error {
	return c.Processor.ProcessPayment(amount)
}

/*
Write a small test for ParkingCheckout using MockPaymentService.

Test both cases:
1. Payment succeeds.
2. Payment fails, and Checkout returns the error.
*/

type MockPaymentService struct {
	ShouldFail bool
}

func (m *MockPaymentService) ProcessPayment(amount float64) error {
	if m.ShouldFail {
		return errors.New("payment failed")
	}
	return nil
}

func DIP() {
	checkoutS := ParkingCheckout{
		Processor: &MockPaymentService{
			ShouldFail: false,
		},
	}

	err := checkoutS.Checkout(65)
	if err != nil {
		fmt.Println("error during checkout")
		return
	}

	checkoutF := ParkingCheckout{
		Processor: &MockPaymentService{
			ShouldFail: true,
		},
	}

	err = checkoutF.Checkout(65)
	if err != nil {
		fmt.Println("error during checkout")
		return
	}
}
