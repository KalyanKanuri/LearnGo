/*
Requirements
1. Create one ParkingLot using []ParkingSpot with spot A101, initially available.
2. Create one ParkingLotPointers with a separate spot B101, also initially available.
3. For the value-slice lot, write FindSpot() using an index and return &lot.Spots[i].
4. For the pointer-slice lot, write FindSpot() returning the matching pointer.
5. Find each spot and mark it occupied through the returned pointer.
6. Print both spots' availability and verify that both original stored spots changed.
*/

package pointers

import "fmt"

type ParkingSpot struct {
	ID       string
	Occupied bool
}
type ParkingLot struct {
	Spots []ParkingSpot
}

func (spot *ParkingSpot) IsAvailable() bool {
	return !spot.Occupied
}

func (lot *ParkingLot) FindSpot(id string) *ParkingSpot {
	for i := range lot.Spots {
		if lot.Spots[i].ID == id {
			return &lot.Spots[i]
		}
	}
	return nil
}

type ParkingLotPointers struct {
	Spots []*ParkingSpot
}

func (lot *ParkingLotPointers) FindSpot(id string) *ParkingSpot {
	for _, spot := range lot.Spots {
		if spot.ID == id {
			return spot
		}
	}
	return nil
}

func occupySpot(spot *ParkingSpot) {
	spot.Occupied = true
}

func PointerObjects() {
	valLot := ParkingLot{
		Spots: make([]ParkingSpot, 0, 2),
	}

	spotA := ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}
	valLot.Spots = append(valLot.Spots, spotA)
	fmt.Printf("Initial A101 Availability Status %t\n", spotA.IsAvailable())

	fSpotA := valLot.FindSpot("A101")
	if fSpotA == nil {
		fmt.Println("Invalid Spot ID Provided")
		return
	}
	occupySpot(fSpotA)

	fmt.Printf("A101 Availability Status after occupy %t\n", spotA.IsAvailable())

	ptrLot := ParkingLotPointers{
		Spots: make([]*ParkingSpot, 0, 2),
	}

	spotB := ParkingSpot{
		ID:       "B101",
		Occupied: false,
	}
	ptrLot.Spots = append(ptrLot.Spots, &spotB)
	fmt.Printf("Initial B101 Availability Status %t\n", spotB.IsAvailable())

	fSpotB := ptrLot.FindSpot("B101")
	if fSpotB == nil {
		fmt.Println("Invalid Spot ID Provided")
		return
	}
	occupySpot(fSpotB)

	fmt.Printf("B101 Availability Status after occupy %t\n", spotB.IsAvailable())
}

/*
Create spotA as a pointer to a new ParkingSpot.
Append that same pointer to lot.Spots.
Change Occupied through lot.Spots[0].
Print the state through spotA and compare the pointers.
*/
func SharedStateDemo() {
	spotA := &ParkingSpot{
		ID:       "A103",
		Occupied: false,
	}
	lot := ParkingLotPointers{
		Spots: make([]*ParkingSpot, 0, 1),
	}
	lot.Spots = append(lot.Spots, spotA)

	fmt.Println("Before Update", spotA.IsAvailable())
	occupySpot(lot.Spots[0])
	fmt.Println("After Update", spotA.IsAvailable())
	fmt.Println("Same Object", lot.Spots[0] == spotA)
}
