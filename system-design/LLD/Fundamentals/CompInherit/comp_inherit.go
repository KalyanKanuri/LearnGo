/*
Implement composition with ParkingLot

Implement this method:
func (lot *ParkingLot) AddSpot(spot *ParkingSpot) error

Requirements:
1. Add the supplied spot to lot.Spots.
2. Return an error if lot is nil.
3. Return an error if spot is nil.
4. Don't add the same spot pointer twice.
5. Don't modify the spot's fields.
*/

package CompInherit

import (
	"errors"
	"fmt"
	parkinglot "fundamentals/ParkingLot"
	"slices"
)

type CILot struct {
	parkinglot.ParkingLot // Embedding
}

func (lot *CILot) AddSpot(spot *parkinglot.ParkingSpot) error {
	if lot == nil {
		return errors.New("invalid lot object")
	}

	if spot == nil {
		return errors.New("invalid spot object")
	}

	if slices.Contains(lot.Spots, spot) {
		return errors.New("spot already exists")
	}

	for _, s := range lot.Spots {
		if s == nil {
			continue
		}

		if s.ID == spot.ID {
			return errors.New("spot already exists")
		}
	}

	lot.Spots = append(lot.Spots, spot)
	return nil
}

/*
Override a promoted method
Suppose ParkingLot already has a FindSpot(id string) method, but CILot needs to log a message before performing the lookup.
Your task: Write a FindSpot(id string) method on CILot that:
1. Prints "Searching in CILot".
2. Calls the embedded ParkingLot's existing FindSpot() method.
3. Returns the resulting *parkinglot.ParkingSpot.
4. Does not duplicate the actual lookup logic.
*/
func (lot *CILot) FindSpot(id string) *parkinglot.ParkingSpot {
	fmt.Println("Searching in CILot")
	return lot.ParkingLot.FindSpot(id)
}
