/* Requirements
1. Define a ParkingSpot struct with ID string and Occupied bool.
2. Add an IsAvailable() method that returns whether the spot is available.
3. Create two spots:
   - A101, initially available.
   - A102, initially occupied.
4. Print each spot's ID and availability.
5. Change A101 to occupied, then print both availability results again.
*/

package object

import "fmt"

type ParkingSpot struct {
	ID       string
	Occupied bool
}

type ParkingLot struct {
	Spots []*ParkingSpot
}

func (lot *ParkingLot) AvailableSpots() int {
	count := 0
	for _, spot := range lot.Spots {
		if !spot.IsAvailable() {
			continue
		}
		count++
	}
	return count
}

func (lot *ParkingLot) FindSpot(id string) *ParkingSpot {
	for _, spot := range lot.Spots {
		if spot.ID == id {
			return spot
		}
	}
	return nil
}

func (ps *ParkingSpot) IsAvailable() bool {
	return !ps.Occupied
}

func occupySpot(spot *ParkingSpot) {
	spot.Occupied = true
}

func CreateObjs() {
	parkingLot := ParkingLot{
		Spots: make([]*ParkingSpot, 0, 3),
	}
	spotA := ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}
	spotB := ParkingSpot{
		ID:       "A102",
		Occupied: true,
	}
	spotC := ParkingSpot{
		ID:       "A103",
		Occupied: false,
	}

	parkingLot.Spots = append(parkingLot.Spots, &spotA, &spotB, &spotC)
	fmt.Printf("Initial available Spots %d\n", parkingLot.AvailableSpots())

	fmt.Printf("A101 availablity status %t\n", spotA.IsAvailable())
	fmt.Printf("A102 availablity status %t\n", spotB.IsAvailable())

	newSpot := parkingLot.FindSpot("A101")
	if newSpot == nil {
		fmt.Println("Invalid spot ID provided")
		return
	}
	occupySpot(newSpot)

	fmt.Printf("A101 availablity status after occupation %t\n", spotA.IsAvailable())
	fmt.Printf("A102 availablity status after occupation %t\n", spotB.IsAvailable())
	fmt.Printf("Final Available Spots %d\n", parkingLot.AvailableSpots())

	// non-existing spot search
	spot := parkingLot.FindSpot("A999")
	if spot == nil {
		fmt.Println("Invalid spot ID provided")
	}
}
