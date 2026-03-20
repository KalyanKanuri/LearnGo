package payrollprocessor

import (
	"fmt"
)

type SalaryPayer interface {
	fmt.Stringer
	calculatePay() float64
}

type SalariedEmployee struct {
	ID           int
	Name         string
	AnnualSalary float64
}

type HourlyEmployee struct {
	ID         int
	Name       string
	HourlyRate float64
	HourlyPay  float64
}

func (salemp SalariedEmployee) calculatePay() float64 {
	return salemp.AnnualSalary / 12
}

func (salemp SalariedEmployee) String() string {
	return fmt.Sprintf("Monthly Pay: %s (Annual: %.2f / 12)", salemp.Name, salemp.AnnualSalary)
}

func (houremp HourlyEmployee) calculatePay() float64 {
	return houremp.HourlyRate * houremp.HourlyPay
}

func (houremp HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly Pay: %s (Hourly: %.2f * %.2f)",
		houremp.Name,
		houremp.HourlyRate,
		houremp.HourlyPay,
	)
}

func printPayDetails[P fmt.Stringer](employee P) {
	fmt.Printf("Processing -- %s\n", employee)
}

func ProcessPayroll(employees []SalaryPayer) {
	fmt.Println("-- Processing Payroll --")
	totalPayroll := 0.0
	for _, employee := range employees {
		printPayDetails(employee)
		pay := employee.calculatePay()
		fmt.Printf("Pay details: %.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Println("Total Monthly payroll:", totalPayroll)
}
