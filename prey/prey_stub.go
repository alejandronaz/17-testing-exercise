package prey

import "testdoubles/positioner"

func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// PreyStub is a stub for the Prey interface
type PreyStub struct {
	// FuncGetSpeed is the function that will be called when GetSpeed is called
	FuncGetSpeed func() (speed float64)
	// FuncGetPosition is the function that will be called when GetPosition is called
	FuncGetPosition func() (position *positioner.Position)
}

func (p *PreyStub) GetSpeed() (speed float64) {
	speed = p.FuncGetSpeed()
	return
}

func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = p.FuncGetPosition()
	return
}
