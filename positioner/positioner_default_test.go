package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("All coordinates are negatives", func(t *testing.T) {
		// arrange
		pos := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}
		to := &positioner.Position{
			X: -1,
			Y: -1,
			Z: -1,
		}

		// act
		linearDistance := pos.GetLinearDistance(from, to)

		// assert
		require.Equal(t, 0.0, linearDistance)
	})
	t.Run("All coordinates are positives", func(t *testing.T) {
		// arrange
		pos := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}
		to := &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}

		// act
		linearDistance := pos.GetLinearDistance(from, to)

		// assert
		require.Equal(t, 0.0, linearDistance)
	})
	t.Run("Coordinates returns lineal distance without decimals", func(t *testing.T) {
		// arrange
		pos := positioner.NewPositionerDefault()
		from := &positioner.Position{
			X: 10,
			Y: 0,
			Z: 0,
		}
		to := &positioner.Position{
			X: 5,
			Y: 0,
			Z: 0,
		}

		// act
		linearDistance := pos.GetLinearDistance(from, to)

		// assert
		require.Equal(t, 5.0, linearDistance)
	})
}
