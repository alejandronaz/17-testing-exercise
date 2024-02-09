package simulator

func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

type CatchSimulatorMock struct {
	// FuncCanCatch is the function that will be called when CanCatch is called
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
	Calls        struct {
		// CanCatch is the number of times CanCatch was called
		CanCatch int
	}
}

func (c *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	// increment the number of times CanCatch was called
	c.Calls.CanCatch++

	// call the function
	canCatch = c.FuncCanCatch(hunter, prey)
	return
}
