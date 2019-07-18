package main

import (
	"fmt"

	p "./PlanesLib"
)

func main() {
	var plane1 = p.NewPlane("Piper", 1992, 320, false)
	var plane2 = p.NewPlane("Cessna", 1972, 260, true)

	fmt.Printf("%+v\n", plane1)

	if plane1 == plane2 {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal!")
	}

	fmt.Println("Plane1 == Plane2: ", plane1.IsEqual(plane2))
}
