package main

import (
	"fundamentals/EncapAbstract"
	object "fundamentals/Objects"
	pointers "fundamentals/Pointers"
	polymorphism "fundamentals/Polymorphism"
	relations "fundamentals/Relations"
	uml "fundamentals/UML"
)

func main() {
	object.CreateObjs()
	pointers.PointerObjects()
	pointers.SharedStateDemo()
	EncapAbstract.EncAbs()
	relations.Rels()
	polymorphism.Polym()
	uml.UMLDemo()
}
