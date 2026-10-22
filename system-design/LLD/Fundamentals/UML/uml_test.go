/*
Test successful occupation
Write one test function named TestOccupySpot_Success that:
1. Creates a UMLLot.
2. Adds one available spot with ID "A101".
3. Calls lot.OccupySpot("A101").
4. Uses t.Fatalf() or t.Errorf() to fail if an error is returned.
5. Verifies the spot is no longer available.
Constraints
- Use Go's standard testing package.
- Don't call UMLDemo() from your test.
- Don't print results using fmt.Println() as your assertions.
- Don't write tests for the other two scenarios yet.
Helpful starting point:
func TestOccupySpot_Success(t *testing.T) {
    // Arrange: create the lot and spot

    // Act: occupy the spot

    // Assert: check the error and resulting state
}
*/

package uml

import (
	parkinglot "fundamentals/ParkingLot"
	"testing"
)

func TestOccupySpot_Success(t *testing.T) {
	// Arrange: create the lot and spot
	lot := UMLLot{}
	spot := parkinglot.ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}
	lot.Spots = append(lot.Spots, &spot)

	// Act: occupy the spot
	err := lot.OccupySpot("A101")

	// Assert: check the error and resulting state
	if err != nil {
		t.Fatalf("Failed to occupy spot: %v\n", err)
	}
	if spot.IsAvailable() {
		t.Errorf("expected spot to be occupied\n")
	}
}

/*
Test an already occupied spot
Add a second test named TestOccupySpot_AlreadyOccupied.

Requirements:
1. Create a lot with a spot "A101" that is already occupied.
2. Call lot.OccupySpot("A101").
3. Verify that an error is returned.
4. Verify that the spot remains occupied.
*/

func TestOccupySpot_AlreadyOccupied(t *testing.T) {
	lot := UMLLot{}
	spot := parkinglot.ParkingSpot{
		ID:       "A101",
		Occupied: true,
	}
	lot.Spots = append(lot.Spots, &spot)

	err := lot.OccupySpot("A101")

	if err == nil {
		t.Fatalf("expected an error when occupying an already occupied spot\n")
	}
	if spot.IsAvailable() {
		t.Errorf("expected spot to be occupied\n")
	}
}

/*
Test a nonexistent spot
Write TestOccupySpot_NotFound.

Requirements:
1. Create a lot containing one spot, "A101".
2. Call lot.OccupySpot("A999").
3. Verify that an error is returned.
4. Verify that the existing "A101" spot remains available.
*/
func TestOccupySpot_NotFound(t *testing.T) {
	lot := UMLLot{}
	spot := parkinglot.ParkingSpot{
		ID:       "A101",
		Occupied: false,
	}
	lot.Spots = append(lot.Spots, &spot)

	err := lot.OccupySpot("A999")
	if err == nil {
		t.Fatalf("expected an error while searching non-existent spot")
	}

	if !spot.IsAvailable() {
		t.Errorf("expected the spot to be available")
	}
}
