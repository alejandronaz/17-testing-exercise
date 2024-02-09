package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSpeed(t *testing.T) {
	t.Run("Return default value", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(0, nil)

		// act
		speed := tuna.GetSpeed()

		// assert
		speedExpected := 0.0
		require.Equal(t, speedExpected, speed)
	})
	t.Run("Return custom value", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(10, nil)

		// act
		speed := tuna.GetSpeed()

		// assert
		speedExpected := 10.0
		require.Equal(t, speedExpected, speed)
	})
}

func TestGetPosition(t *testing.T) {
	t.Run("Return default value", func(t *testing.T) {
		// arrange
		tuna := prey.NewTuna(0, &positioner.Position{
			X: 0,
			Y: 0,
			Z: 0,
		})

		// act
		position := tuna.GetPosition()

		// assert
		positionExpected := &positioner.Position{
			X: 0,
			Y: 0,
			Z: 0,
		}
		require.Equal(t, positionExpected, position)
	})
	t.Run("Return custom value", func(t *testing.T) {
		// arrange
		pos := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		tuna := prey.NewTuna(0, pos)

		// act
		position := tuna.GetPosition()

		// assert
		positionExpected := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		require.Equal(t, positionExpected, position)
	})
}
