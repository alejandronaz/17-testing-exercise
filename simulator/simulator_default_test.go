package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCatch(t *testing.T) {
	t.Run("Can catch succes", func(t *testing.T) {
		// arrange

		// positioner stub
		pos := positioner.NewPositionerStub()
		pos.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 10.0
		}

		// catch simulator
		// - max time to catch
		maxTimeToCatch := 15.0
		sim := simulator.NewCatchSimulatorDefault(maxTimeToCatch, pos)

		// - hunter and prey
		hunter := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 10.0,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 5.0,
		}

		// act
		ok := sim.CanCatch(hunter, prey)

		// assert
		require.True(t, ok)
	})
	t.Run("Can't catch success because hunter is slower than prey", func(t *testing.T) {
		// arrange

		// positioner stub
		pos := positioner.NewPositionerStub()
		pos.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 10.0
		}

		// catch simulator
		// - max time to catch
		maxTimeToCatch := 15.0
		sim := simulator.NewCatchSimulatorDefault(maxTimeToCatch, pos)

		// - hunter and prey
		hunter := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 5.0,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 10.0,
		}

		// act
		ok := sim.CanCatch(hunter, prey)

		// assert
		require.False(t, ok)
	})
	t.Run("Can't catch success beacuse there is no time to catch", func(t *testing.T) {
		// arrange

		// positioner stub
		pos := positioner.NewPositionerStub()
		pos.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 100.0
		}

		// catch simulator
		// - max time to catch
		maxTimeToCatch := 15.0
		sim := simulator.NewCatchSimulatorDefault(maxTimeToCatch, pos)

		// - hunter and prey
		hunter := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 10.0,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed: 5.0,
		}

		// act
		ok := sim.CanCatch(hunter, prey)

		// assert
		require.False(t, ok)
	})
}
