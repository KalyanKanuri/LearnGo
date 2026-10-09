/*
Fix the design in flaw_code.go, everything works fine in flaw_code.go but we should follow
the SRP and redesign the code.

-- REQUIREMENTS:
> ParkingLot should not calculate Fee.
> Introduce seperate PricingService responsible for fee calculation.
> Keep ParkingTicket focused on representing the parking session.

-- DELIVERABLES:
> Write refactored go structs and methods
> Show how the components are constructed and used in main()
> Write table driven tests for fee calculation
	a. session shorter than one hour
	b. session longer than one hour
*/

package main

import (
	"fmt"
	"time"
)

type ParkingLotFixed struct{}

type ParkingTicketFixed struct {
	TicketID  uint
	EntryTime time.Time
	ExitTime  time.Time
}

func NewParkingTicket(ticketID uint, entryTime, exitTime time.Time) *ParkingTicketFixed {
	return &ParkingTicketFixed{
		TicketID:  ticketID,
		EntryTime: entryTime,
		ExitTime:  exitTime,
	}
}

type PricingService struct {
	HourlyRate float64
}

func NewPricingService(hourlyRate float64) *PricingService {
	return &PricingService{
		HourlyRate: hourlyRate,
	}
}

func (ps *PricingService) CalculateFee(t *ParkingTicketFixed) float64 {
	hours := max(int(t.ExitTime.Sub(t.EntryTime).Hours()), 1)
	return ps.HourlyRate * float64(hours)
}

func (ps *PricingService) DisplayFee(t *ParkingTicketFixed)

func main() {
	tkt := NewParkingTicket(1, time.Now(), time.Now().Add(1*time.Hour))
	ps := NewPricingService(10.5)
	fee := ps.CalculateFee(tkt)
	fmt.Println(fee)
}
