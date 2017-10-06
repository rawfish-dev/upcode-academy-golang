package main

import (
	"fmt"
)

var _ (Vehicle) = new(QualityVehicle)
var _ (Vehicle) = new(LousyVehicle)

type Vehicle interface {
	Transport()
}

type QualityVehicle struct {
	Name string
}

func (q QualityVehicle) Transport() {
	fmt.Println(q.Name, "is a great vehicle and the engine purrs.")
}

type LousyVehicle struct {
	Name string
}

func (l LousyVehicle) Transport() {
	fmt.Println(l.Name, "is making some strange noises when shifting gears.")
}
