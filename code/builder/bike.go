package builder

type Bike struct {
	v Product
}

func (b *Bike) SetWheels() Builder {
	b.v.Wheels = 2
	return b
}

func (b *Bike) SetSeats() Builder {
	b.v.Seats = 2
	return b
}

func (b *Bike) SetStructure() Builder {
	b.v.Structure = "Motorbike"
	return b
}

func (b *Bike) GetVehicle() Product {
	return b.v
}
