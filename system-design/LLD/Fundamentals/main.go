package main

import (
	"fundamentals/EncapAbstract"
	object "fundamentals/Objects"
	pointers "fundamentals/Pointers"
)

func main() {
	object.CreateObjs()
	pointers.PointerObjects()
	pointers.SharedStateDemo()
	EncapAbstract.EncAbs()
}
