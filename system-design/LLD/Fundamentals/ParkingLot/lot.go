package parkinglot

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

func (lot *ParkingLot) FindSpot(id string) *ParkingSpot {
	for i := range lot.Spots {
		if lot.Spots[i].ID == id {
			return lot.Spots[i]
		}
	}
	return nil
}
