package dip

import (
	"testing"
)

func TestParkingCheckout(t *testing.T) {
	t.Run("Checkout Success", func(t *testing.T) {
		checkoutS := ParkingCheckout{
			Processor: &MockPaymentService{
				ShouldFail: false,
			},
		}

		err := checkoutS.Checkout(65)
		if err != nil {
			t.Errorf("error during checkout %v", err)
		}
	})

	t.Run("Checkout Failure", func(t *testing.T) {
		checkoutF := ParkingCheckout{
			Processor: &MockPaymentService{
				ShouldFail: true,
			},
		}

		err := checkoutF.Checkout(65)
		if err == nil {
			t.Errorf("exepected checkout to fail, got nil error")
		}
	})
}
