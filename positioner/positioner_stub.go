package positioner

func NewPositionerStub() *PositionerStub {
	return &PositionerStub{}
}

// PositionerStub is a stub for the Positioner interface
type PositionerStub struct {
	// FuncGetLinearDistance is the function that will be called when GetLinearDistance is called
	// - externalize the implementation of the function to allow the test to control the return value
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	// call the function
	linearDistance = p.FuncGetLinearDistance(from, to)
	return
}
