package PlanesLib

import (
	"fmt"
)

func (LHS *Airplane) IsEqual(RHS *Airplane) bool {
	return LHS.Model == RHS.Model && LHS.Year == RHS.Year && LHS.Hp == RHS.Hp && LHS.ForSale == RHS.ForSale
}

func (plane *Airplane) String() string {
	return fmt.Sprintf("Model: %s\n Year: %d\n Hp: %d\n For Sale: %t", plane.Model, plane.Year, plane.Hp, plane.ForSale)
}
