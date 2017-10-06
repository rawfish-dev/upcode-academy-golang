package main

import (
	"fmt"
)

/*
Drivers and Vehicles
Let’s write a simple program that leverages interfaces to let drivers drive vehicles.
This will allow us to see why interfaces are useful and practical in reality.
*/

// 1) Write an interface for the driver
// 2) Write an interface for the vehicle
// 3) Create the structs that satisfy either interface

func main() {
	david := CapableDriver{"David"}
	bob := PoorDriver{"Bob"}

	mercedes := QualityVehicle{"Mercedes"}
	armyJeep := LousyVehicle{"Army Jeep"}

	david.Drive(mercedes)
	PrintSeparator()

	david.Drive(armyJeep)
	PrintSeparator()

	bob.Drive(mercedes)
	PrintSeparator()

	bob.Drive(armyJeep)
	PrintSeparator()
}

func PrintSeparator() {
	fmt.Printf("-------\n\n")
}
