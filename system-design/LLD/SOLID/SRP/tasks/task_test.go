package main

import (
	"testing"
	"time"
)

type TestFeeInstance struct {
	name     string
	duration time.Duration
	rate     float64
	expected float64
}

func TestPricingService_CalculateFee(t *testing.T) {
	tests := []TestFeeInstance{
		{
			name:     "Under One Hour",
			duration: 45 * time.Minute,
			rate:     10.5,
			expected: 10.5,
		},
		{
			name:     "One Hour",
			duration: time.Hour,
			rate:     10.5,
			expected: 10.5,
		},
		{
			name:     "Above 2 Hours",
			duration: 2*time.Hour + 30*time.Minute,
			rate:     10.5,
			expected: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := time.Now()
			ticket := NewParkingTicket(1, entry, entry.Add(tt.duration))
			service := NewPricingService(tt.rate)

			result := service.CalculateFee(ticket)

			if result != tt.expected {
				t.Errorf("CalculateFee() got: %.2f, want : %.2f\n", result, tt.expected)
			}
		})

	}
}
