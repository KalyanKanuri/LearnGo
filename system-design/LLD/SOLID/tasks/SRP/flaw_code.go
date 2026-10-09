package main

import (
	"fmt"
	"time"
)

type ParkingTicket struct {
	LicensePlate string
	EntryTime    time.Time
	ExitTime     time.Time
}

type ParkingLot struct {
	HourlyRate float64
}

func (p *ParkingLot) CalculateFee(ticket ParkingTicket) float64 {
	hours := max(int(ticket.ExitTime.Sub(ticket.EntryTime).Hours()), 1)
	return p.HourlyRate * float64(hours)
}

func (p *ParkingLot) PrintTicket(ticket ParkingTicket) {
	fmt.Println("License Plate: ", ticket.LicensePlate)
	fmt.Println("Fee: ", p.CalculateFee(ticket))
}
