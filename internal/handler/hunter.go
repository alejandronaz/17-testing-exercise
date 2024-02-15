package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// get the bytes from the request body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error reading the request body"))
			return
		}

		// convert the bytes to a map
		var bodyMap map[string]any
		if json.Unmarshal(bodyBytes, &bodyMap) != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("json is malformed"))
			return
		}

		// validate the keys
		if !validateKeyExistance(bodyMap, "speed", "position") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("keys are missing"))
			return
		}

		// convert the map to a struct
		var body RequestBodyConfigPrey
		if json.Unmarshal(bodyBytes, &body) != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("json types are wrong"))
			return
		}

		// process
		h.pr = prey.NewTuna(body.Speed, body.Position)

		// response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("prey configured"))
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the bytes from the request body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error reading the request body"))
			return
		}

		// convert the bytes to a map
		var bodyMap map[string]any
		if json.Unmarshal(bodyBytes, &bodyMap) != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("json is malformed"))
			return
		}

		// validate the keys
		if !validateKeyExistance(bodyMap, "speed", "position") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("keys are missing"))
			return
		}

		// convert the map to a struct
		var body RequestBodyConfigHunter
		if json.Unmarshal(bodyBytes, &body) != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("json types are wrong"))
			return
		}

		// process
		// - simulator
		cfgSim := simulator.ConfigCatchSimulatorDefault{
			MaxTimeToCatch: 15.0,
			Positioner:     positioner.NewPositionerDefault(),
		}
		sim := simulator.NewCatchSimulatorDefault(&cfgSim)
		// - hunter
		h.ht = hunter.NewWhiteShark(hunter.ConfigWhiteShark{
			Speed:     body.Speed,
			Position:  body.Position,
			Simulator: sim,
		})

		// response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hunter configured"))
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// process
		if h.ht == nil || h.pr == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("hunter or prey not configured"))
			return
		}

		duration, err := h.ht.Hunt(h.pr)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("could not hunt"))
			return
		}

		// response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("hunt done in %.2f", duration)))
	}
}

func validateKeyExistance(m map[string]any, keys ...string) bool {
	for _, key := range keys {
		_, ok := m[key]
		if !ok {
			return false
		}
	}
	return true
}
