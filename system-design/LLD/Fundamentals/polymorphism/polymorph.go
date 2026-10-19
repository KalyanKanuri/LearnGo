/*
Requirements
 1. Define a Vehicle interface with this method:
    CanFit(spot *ParkingSpot) bool
 2. Define a Car type.
 3. Implement CanFit() for Car using a pointer receiver or value receiver—choose one and be consistent.
 4. For this exercise, assume a car can fit only in a spot whose ID is 101. Return true for that spot and false otherwise.
 5. Don't add type switches, reflection, or a ParkingService yet.
*/
package polymorphism

import (
	"fmt"
	parkinglot "fundamentals/ParkingLot"
)

type Vehicle interface {
	CanFit(spot *parkinglot.ParkingSpot) bool
}

type Car struct{}

func (c *Car) CanFit(spot *parkinglot.ParkingSpot) bool {
	if spot == nil {
		fmt.Println("Invalid Spot")
		return false
	}
	if spot.ID == "101" {
		return true
	}
	return false
}

/*
Without changing Car, write a small function that accepts the Vehicle interface and calls CanFit().
Use this signature:
func CheckFit(v Vehicle, spot *parkinglot.ParkingSpot) bool


Requirements:
1. Accept any Vehicle, not specifically a *Car.
2. Call v.CanFit(spot) and return the result.
3. Don't use type assertions or type switches.
4. Handle a nil Vehicle interface safely if you can—think carefully about what v == nil checks.
*/

func CheckFit(v Vehicle, spot *parkinglot.ParkingSpot) bool {
	if v == nil {
		return false
	}
	return v.CanFit(spot)
}

/*
Prove that a different type works
Now implement a Motorcycle type that also satisfies Vehicle.
Requirements:
1. Define Motorcycle.
2. Implement CanFit(spot *parkinglot.ParkingSpot) bool.
3. For this exercise, a motorcycle can fit only when the spot ID is "102".
4. Handle a nil spot safely.
5. Don't modify Vehicle or CheckFit().
*/

type Motorcycle struct{}

func (m *Motorcycle) CanFit(spot *parkinglot.ParkingSpot) bool {
	if spot == nil {
		fmt.Println("Invalid Spot")
		return false
	}
	if spot.ID == "102" {
		return true
	}
	return false
}

/*
Demonstrate polymorphism
You already have Vehicle, Car, Motorcycle, and CheckFit(). Now write a small test function to prove they work together.
Requirements
1. Create a ParkingSpot with ID "101" and another with ID "102".
2. Create a *Car and a *Motorcycle.
3. Call CheckFit() for both vehicles against both spots.
4. Print the four results.
5. Don't change the existing implementations.
*/
func Polym() {
	car := Car{}
	spot1 := parkinglot.ParkingSpot{
		ID:       "101",
		Occupied: false,
	}

	motorcycle := Motorcycle{}
	spot2 := parkinglot.ParkingSpot{
		ID:       "102",
		Occupied: false,
	}

	fmt.Println(CheckFit(&car, &spot1))
	fmt.Println(CheckFit(&car, &spot2))
	fmt.Println(CheckFit(&motorcycle, &spot1))
	fmt.Println(CheckFit(&motorcycle, &spot2))
}
