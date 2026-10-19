package main

import (
	"fundamentals/EncapAbstract"
	object "fundamentals/Objects"
	pointers "fundamentals/Pointers"
	relations "fundamentals/Relations"
)

func main() {
	object.CreateObjs()
	pointers.PointerObjects()
	pointers.SharedStateDemo()
	EncapAbstract.EncAbs()
	relations.Rels()
}
