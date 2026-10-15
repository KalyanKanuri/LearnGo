/*
Refactor ParkingSpot so it controls its own state.
Requirements
1. Make the internal occupied state unexported.
2. Keep the spot ID readable from outside the package, but don't allow callers to change it directly.
3. Implement IsAvailable() bool to report whether the spot is available.
4. Implement Occupy() error that:
   - Occupies an available spot.
   - Returns an error if the spot is already occupied.
   - Leaves the state unchanged when the operation fails.
5. Add a Release() error method that releases an occupied spot and returns an error if it is already available.
Constraints: Don't add a ParkingLot or interfaces yet. Focus only on one ParkingSpot.
*/

package EncapAbstract

import (
	"errors"
	"fmt"
)

type ParkingSpot struct {
	id       int
	occupied bool // keep it private
}

func (spot *ParkingSpot) ID() int {
	if spot == nil {
		fmt.Println("Invalid spot object")
		return 0
	}
	return spot.id
}

func (spot *ParkingSpot) IsAvailable() bool {
	if spot == nil {
		fmt.Println("Invalid spot object")
		return false
	}
	return !spot.occupied
}

/*
Requirements:
1. Modify only the Occupy() method.
2. Use IsAvailable() to determine whether the spot can be occupied.
3. Preserve the current behavior: return an error if the spot is already occupied.
4. Change the state only when occupation succeeds.
5. Keep the nil-receiver check.
*/

func (spot *ParkingSpot) Occupy() error {
	if spot == nil {
		return errors.New("Invalid spot object")
	}

	if !spot.IsAvailable() {
		return errors.New("Spot is already occupied")
	}
	spot.occupied = true
	return nil
}

func (spot *ParkingSpot) Release() error {
	if spot == nil {
		return errors.New("Invalid spot object")
	}

	if spot.occupied == false {
		return errors.New("Spot is already available")
	}
	spot.occupied = false
	return nil
}

func EncAbs() {
	spot := ParkingSpot{
		id:       101,
		occupied: false,
	}

	err := spot.Occupy()
	if err != nil {
		fmt.Println(err)
	}

	err = spot.Occupy()
	if err != nil {
		fmt.Println(err)
	}

	err = spot.Release()
	if err != nil {
		fmt.Println(err)
	}

	err = spot.Release()
	if err != nil {
		fmt.Println(err)
	}
}
