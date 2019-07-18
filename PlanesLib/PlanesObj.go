package PlanesLib

type Airplane struct {
	Model   string // ""
	Year    int    // 0
	Hp      int    // 0
	ForSale bool   // false
}

func NewPlaneDefault() *Airplane {
	return new(Airplane) // new automatically returns a pointer to the object
}

func NewPlane(model string, year int, hp int, sale bool) *Airplane {
	newPlane := new(Airplane)

	newPlane.Model = model
	newPlane.Year = year
	newPlane.Hp = hp
	newPlane.ForSale = sale

	return newPlane
}
