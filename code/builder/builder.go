package builder

// Builder is the build process of a product.
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
}

// Product defines the construct of a product.
type Product struct {
	Wheels int
	Seats int
	Structure string
}

// ManufacturingDirector manages the process to build a product.
type ManufacturingDirector struct {
	builder Builder
}

// Construct defines the sequence of build process.
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder is used to set a builder for different products.
func (f *ManufacturingDirector) SetBuilder(b Builder) {
	f.builder = b
}
