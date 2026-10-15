# Single Responsibility Principle

In Solid, S - stands for Single Responsibility Principle that means we should ensure that an object should hold responsibilities only related to the object

An Object should have it’s own responsibilities, properties and behavior and these should be related.

This doesn’t mean, a class/Struct should have only one method, it can have multiple methods but those methods should be related to the Object.

e.g.,

```go
// Valid Scenario

type ParkingFloor struct {
 spot []ParkingSpot
}

func (ps *ParkingFloor) AddSpot(s ParkingSpot)  {}
func (ps *ParkingFloor) RemoveSpot(s ParkingSpot) {}
func (ps *ParkingFloor) AvailableSpots() []ParkingSpot {}
func (ps *ParkingFloor) OccupiedSpots() []ParkingSpot {}
```

```go
// Invalid Scenario

// to the same ParkingFloor defined above we should not add below methods

func (ps *ParkingFloor) CalculateFee() {} // this should belong to Payment
func (ps *ParkingFloor) NorifyUsers() {}  // this should belong to Notifier
func (ps *ParkingFloor) GenReports() {}   // this should belogn to Reporter
```

so SRP isn’t about number of methods it’s about the responsibility.

IDENTIFICATION:
whenever we design a class we need to identify if the method changes in future then class also needs to be changed or not if it’s not required to change the class then the method belongs somewhere else.

suppose if a parking spot is marked as VIP tomorrow then the vehicle also needs to be changed? no hence the parking spot responsibility should not be written in vehicle
