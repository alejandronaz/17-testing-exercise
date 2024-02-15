package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testdoubles/internal/handler"
	"testdoubles/internal/hunter"
	"testdoubles/internal/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigureHunter(t *testing.T) {
	t.Run("Hunter is configured successfully", func(t *testing.T) {
		// arrange
		// - hunter
		hunter := hunter.NewHunterMock()
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hunter, prey)
		// - request
		req := httptest.NewRequest(
			"POST",
			"/hunter/configure-hunter",
			strings.NewReader(`{ "speed": 30, "position": {"X": 1, "Y": 1, "Z": 1} }`),
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.ConfigureHunter()
		hdFunc(res, req)

		// assert
		expectedBody := "hunter configured"
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, http.StatusOK, res.Code)
	})
	t.Run("Hunter is not configured successfully - missing body fields", func(t *testing.T) {
		// arrange
		// - hunter
		hunter := hunter.NewHunterMock()
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hunter, prey)
		// - request
		req := httptest.NewRequest(
			"POST",
			"/hunter/configure-hunter",
			strings.NewReader(`{ "speed": 30 }`),
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.ConfigureHunter()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestConfigurePrey(t *testing.T) {
	t.Run("Prey is configured successfully", func(t *testing.T) {
		// arrange
		// - hunter
		hunter := hunter.NewHunterMock()
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hunter, prey)
		// - request
		req := httptest.NewRequest(
			"POST",
			"/hunter/configure-prey",
			strings.NewReader(`{ "speed": 10, "position": {"X": 1, "Y": 1, "Z": 1} }`),
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.ConfigurePrey()
		hdFunc(res, req)

		// assert
		expectedBody := "prey configured"
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, http.StatusOK, res.Code)
	})
	t.Run("Prey is not configured successfully - missing body fields", func(t *testing.T) {
		// arrange
		// - hunter
		hunter := hunter.NewHunterMock()
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hunter, prey)
		// - request
		req := httptest.NewRequest(
			"POST",
			"/hunter/configure-prey",
			strings.NewReader(`{ "speed": 30 }`),
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.ConfigurePrey()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestHunt(t *testing.T) {
	t.Run("Hunt is successful", func(t *testing.T) {
		// arrange

		// - configure
		// - hunter
		hunter := hunter.NewHunterMock()
		hunter.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 10.0, nil
		}
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hunter, prey)

		// - request to hunt
		req := httptest.NewRequest(
			"POST",
			"/hunter/hunt",
			nil,
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.Hunt()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, hunter.Calls.Hunt, 1)
		require.Equal(t, "hunt done in 10.00", res.Body.String())
	})

	t.Run("Hunt is not successful - prey not configured", func(t *testing.T) {
		// arrange

		// - configure
		// - hunter
		hunter := hunter.NewHunterMock()
		hunter.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 10.0, nil
		}
		// - handler
		hd := handler.NewHunter(hunter, nil)

		// - request to hunt
		req := httptest.NewRequest(
			"POST",
			"/hunter/hunt",
			nil,
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.Hunt()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("Hunt is not successful - hunter not catch prey", func(t *testing.T) {
		// arrange

		// - configure
		// - hunter
		hun := hunter.NewHunterMock()
		hun.HuntFunc = func(pr prey.Prey) (duration float64, err error) {
			return 0.0, hunter.ErrCanNotHunt
		}
		// - prey
		prey := prey.NewPreyStub()
		// - handler
		hd := handler.NewHunter(hun, prey)

		// - request to hunt
		req := httptest.NewRequest(
			"POST",
			"/hunter/hunt",
			nil,
		)
		// - response
		res := httptest.NewRecorder()

		// act
		hdFunc := hd.Hunt()
		hdFunc(res, req)

		// assert
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "could not hunt", res.Body.String())
	})
}
