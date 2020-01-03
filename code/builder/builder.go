package builder

type buildProcess interface {
	setWheels() buildProcess
	setSeats() buildProcess
	setStructure() buildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) setWheels() buildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) setSeats() buildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) setStructure() buildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) setWheels() buildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) setSeats() buildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) setStructure() buildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type ManufacturingDirector struct {
	builder buildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.setSeats().setStructure().setWheels()
}

func (f *ManufacturingDirector) SetBuilder(b buildProcess) {
	f.builder = b
}
