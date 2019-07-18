package main

import (
	"fmt"

	p "./PlanesLib"
)

func main() {
	var plane1 = p.NewPlane("Piper", 1992, 320, false)
	var plane2 = p.NewPlane("Cessna", 1972, 260, true)
	var plane3 = p.NewPlane("Cessna", 1972, 260, true)
	var plane4 = p.NewPlane("Cessna", 1972, 300, true)
	var plane5 = p.NewPlaneDefault()

	fmt.Printf("%+v\n", plane1)

	if plane1 == plane2 {
		fmt.Println("Planes 1 & 2 are the same object (Address)")
	} else {
		fmt.Println("Planes 1 & 2 are not the same object (Address)!")
	}

	if plane2 == plane3 {
		fmt.Println("Planes 2 & 3 are the same object (Address)")
	} else {
		fmt.Println("Planes 2 & 3 are not the same object (Address)!")
	}

	fmt.Println("Plane2 == Plane3: ", plane2.IsEqual(plane3))
	fmt.Println("Plane4 == Plane5: ", plane4.IsEqual(plane5))
}
