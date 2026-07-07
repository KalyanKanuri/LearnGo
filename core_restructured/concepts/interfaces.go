package concepts

import (
	"fmt"
)

/*
  Interfaces:
  ----------
	> An Interface is simply a collection of method signatures
	> Any type that satisfies the interface method signatures is treated as an interface
*/

type Vehicle interface {
	startVehicle()
	stopVehicle()
}

// here Bike struct implements all the methods required by the Vehicle interface so *Bike satisfies the Vehicle interface
type Bike struct {
	BikeName    string
	EngineModel string
	MaxSpeed    int
}

func (b *Bike) startVehicle() {
	fmt.Println(b.BikeName, "Started")

}

func (b *Bike) stopVehicle() {
	fmt.Println(b.BikeName, "Stopped")
}

func operateVehicle(v Vehicle) {
	v.startVehicle()
	v.stopVehicle()
}

func InterfacesInGo() {
	xtreme_100 := &Bike{
		BikeName:    "Xtreme 100",
		EngineModel: "4V",
		MaxSpeed:    160,
	}
	operateVehicle(xtreme_100)
}
