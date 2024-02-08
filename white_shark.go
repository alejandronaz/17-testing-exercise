package hunt

import (
	"errors"
)

var (
	// ErrSharkIsTired indicates that the shark is tired
	ErrSharkIsTired = errors.New("can not hunt, shark is tired")
	// ErrSharkIsNotHungry indicates that the shark is not hungry
	ErrSharkIsNotHungry = errors.New("can not hunt, shark is not hungry")
	// ErrSharkIsSlower indicates that the shark could not catch the prey
	ErrSharkIsSlower = errors.New("can not hunt, shark is slower than the prey")
	// ErrSharkThereIsNoTuna indicates that the shark could not catch the prey
	ErrSharkThereIsNoTuna = errors.New("can not hunt, there is no tuna")
)

// NewWhiteShark creates a new WhiteShark
func NewWhiteShark(hungry bool, tired bool, speed float64) (w *WhiteShark) {
	w = &WhiteShark{
		Hungry: hungry,
		Tired:  tired,
		Speed:  speed,
	}
	return
}

// WhiteShark is an implementation of the Hunter interface
type WhiteShark struct {
	// Hungry indicates if the shark is hungry
	Hungry bool
	// Tired indicates if the shark is tired
	Tired bool
	// Speed indicates the speed of the shark
	Speed float64
}

func (w *WhiteShark) Hunt(tuna *Tuna) (err error) {
	// check if there is a tuna
	if tuna == nil {
		err = ErrSharkThereIsNoTuna
		return
	}

	// check if the shark can hunt
	if !w.Hungry {
		err = ErrSharkIsNotHungry
		return
	}
	if w.Tired {
		err = ErrSharkIsTired
		return
	}
	if w.Speed < tuna.Speed {
		err = ErrSharkIsSlower
		return
	}

	// hunt done
	w.Hungry = false
	w.Tired = true
	return
}
