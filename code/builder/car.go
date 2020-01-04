package builder

type Car struct {
	v Product
}

func (c *Car) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}

func (c *Car) SetSeats() Builder {
	c.v.Seats = 5
	return c
}

func (c *Car) SetStructure() Builder {
	c.v.Structure = "Car"
	return c
}

func (c *Car) GetVehicle() Product {
	return c.v
}