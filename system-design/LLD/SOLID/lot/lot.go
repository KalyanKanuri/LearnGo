package parkinglot

import "errors"

type ParkingSpot struct {
	ID       string
	Occupied bool
}

type ParkingLot struct {
	Spots []*ParkingSpot
}

func (spot *ParkingSpot) IsAvailable() bool {
	return !spot.Occupied
}

func (spot *ParkingSpot) Occupy() error {
	if spot == nil {
		return errors.New("Invalid spot object")
	}

	if !spot.IsAvailable() {
		return errors.New("Spot is already occupied")
	}
	spot.Occupied = true
	return nil
}

func (lot *ParkingLot) FindSpot(id string) *ParkingSpot {
	for i := range lot.Spots {
		if lot.Spots[i].ID == id {
			return lot.Spots[i]
		}
	}
	return nil
}
