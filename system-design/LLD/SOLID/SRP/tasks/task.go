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

// TODO: Implement floor management
// keep the struct for future development
type ParkingLot struct{}

type ParkingTicket struct {
	TicketID  uint
	EntryTime time.Time
	ExitTime  time.Time
}

func NewParkingTicket(ticketID uint, entryTime, exitTime time.Time) *ParkingTicket {
	return &ParkingTicket{
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

func (ps *PricingService) CalculateFee(t *ParkingTicket) float64 {
	hours := int(t.ExitTime.Sub(t.EntryTime).Hours())
	if hours < 1 {
		hours = 1
	}
	return ps.HourlyRate * float64(hours)
}

func main() {
	now := time.Now()
	tkt := NewParkingTicket(1, now, now.Add(1*time.Hour))
	ps := NewPricingService(10.5)
	fee := ps.CalculateFee(tkt)
	fmt.Println(fee)
}
