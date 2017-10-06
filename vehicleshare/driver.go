package main

import (
	"fmt"
)

var _ (Driver) = new(CapableDriver)
var _ (Driver) = new(PoorDriver)

type Driver interface {
	Drive(Vehicle)
}

type CapableDriver struct {
	Name string
}

func (c CapableDriver) Drive(vehicle Vehicle) {
	fmt.Println(c.Name, "is driving very well!")
	vehicle.Transport()
}

type PoorDriver struct {
	Name string
}

func (p PoorDriver) Drive(vehicle Vehicle) {
	fmt.Println(p.Name, "is driving pretty dangerously!")
	vehicle.Transport()
}
