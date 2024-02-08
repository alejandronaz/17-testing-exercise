package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrange
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("Atun", 5.0)

		// act
		err := shark.Hunt(tuna)

		// assert
		require.NoError(t, err)
		require.False(t, shark.Hungry)
		require.True(t, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrange
		shark := hunt.NewWhiteShark(false, false, 10.0)
		tuna := hunt.NewTuna("Atun", 5.0)
		errExpected := hunt.ErrSharkIsNotHungry

		// act
		err := shark.Hunt(tuna)

		// assert
		require.Equal(t, errExpected, err)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrange
		shark := hunt.NewWhiteShark(true, true, 10.0)
		tuna := hunt.NewTuna("Atun", 5.0)
		errExpected := hunt.ErrSharkIsTired

		// act
		err := shark.Hunt(tuna)

		// assert
		require.Equal(t, errExpected, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrange
		shark := hunt.NewWhiteShark(true, false, 10.0)
		tuna := hunt.NewTuna("Atun", 15.0)
		errExpected := hunt.ErrSharkIsSlower

		// act
		err := shark.Hunt(tuna)

		// assert
		require.Equal(t, errExpected, err)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrange
		shark := hunt.NewWhiteShark(true, false, 10.0)
		var tuna *hunt.Tuna
		errExpected := hunt.ErrSharkThereIsNoTuna

		// act
		err := shark.Hunt(tuna)

		// assert
		require.Equal(t, errExpected, err)
	})
}
