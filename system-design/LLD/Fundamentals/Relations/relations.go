/*
Requirements:
1. Find the spot with the given ID.
2. Remove its pointer from lot.spots.
3. Return true if a spot was removed; otherwise, return false.
4. Do not modify the ParkingSpot object's fields.
5. If another variable references the removed spot, that variable should still be able to access the object.
*/

package relations

import (
	"fmt"
	parkinglot "fundamentals/ParkingLot"
	"slices"
)

// Aggregation - An object can manage other independent objects.
// Association - An object can work with other object.
/*
type ParkingAttendant struct {
	ParkingLot *ParkingLot // here parking lot is associated with ParkingAttendant
}
*/
// Composition - An object can manage other dependent objects

// this is Composition
type RelLot struct {
	parkinglot.ParkingLot
}

// here even though if we remove spot pointer from spots, the spot object will exist independently.
func (lot *RelLot) RemoveSpot(id string) bool {
	if id == "" {
		fmt.Println("Invalid ID")
		return false
	}

	spot := lot.FindSpot(id)
	if spot == nil {
		fmt.Println("Invalid Spot")
		return false
	}

	for i, s := range lot.Spots {
		if s == spot {
			lot.Spots = slices.Delete(lot.Spots, i, i+1)
			return true
		}
	}
	return false
}

/*
Requirements
1. Define a ParkingService struct.
2. Implement this method:
   func (s *ParkingService) FindSpot(
       lot *ParkingLot,
       id string,
   ) *ParkingSpot
3. The method must use lot.FindSpot(id) and return its result.
4. If lot is nil, return nil rather than panic.
5. Don't store lot as a field in ParkingService.
*/

type ParkingService struct{}

func (s *ParkingService) FindSpot(
	lot *parkinglot.ParkingLot,
	id string,
) *parkinglot.ParkingSpot {
	if lot == nil {
		return nil
	}
	return lot.FindSpot(id)
}

func Rels() {
	lot := parkinglot.ParkingLot{}

	spot1 := parkinglot.ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}

	spot2 := parkinglot.ParkingSpot{
		ID:       "A102",
		Occupied: false,
	}
	lot.Spots = append(lot.Spots, &spot1, &spot2)

	rLot := RelLot{
		ParkingLot: lot,
	}

	result := rLot.RemoveSpot("A101")
	fmt.Println(result)
}
