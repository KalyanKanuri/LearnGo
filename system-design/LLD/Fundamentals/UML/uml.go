/*
Implement ParkingLot.OccupySpot(id string) error.
It should:
1. Find the spot using its ID.
2. Return an error if the spot doesn't exist.
3. Delegate the state transition to ParkingSpot.Occupy().
4. Return any error from the spot unchanged.
5. Avoid directly modifying the spot's private occupancy state.
*/

package uml

import (
	"errors"
	"fmt"
	parkinglot "fundamentals/ParkingLot"
)

type UMLLot struct {
	parkinglot.ParkingLot
}

func (lot *UMLLot) OccupySpot(id string) error {
	if lot == nil {
		return errors.New("invalid lot object")
	}

	if id == "" {
		return errors.New("invalid spot id")
	}

	spot := lot.FindSpot(id)
	if spot == nil {
		return errors.New("spot doesn't exist")
	}

	return spot.Occupy()
}

func UMLDemo() {
	lot := UMLLot{}

	// An existing available spot becomes occupied.
	spot1 := parkinglot.ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}
	lot.Spots = append(lot.Spots, &spot1)

	err := lot.OccupySpot("A101")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Case1: status", spot1.IsAvailable())

	// An already occupied spot returns an error without changing state
	err = lot.OccupySpot("A101")
	if err == nil {
		fmt.Println("test failed as error is expected")
		return
	}

	if spot1.IsAvailable() == true {
		fmt.Println("failed unexpected spot state")
		return
	}
	fmt.Println("Case2 passed: error returned and spot remains occupied")

	// A nonexistent spot ID returns an error.
	err = lot.OccupySpot("A102")
	if err == nil {
		fmt.Println("test failed as error is expected")
		return
	}

	if spot1.IsAvailable() == true {
		fmt.Println("failed unexpected spot state")
		return
	}
	fmt.Println("Case3 passed: error returned and spot remains occupied")
}
