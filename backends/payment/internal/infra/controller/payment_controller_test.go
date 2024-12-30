package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	di "github.com.br/gibranct/payment/internal/infra/DI"
	"github.com.br/gibranct/payment/internal/infra/controller"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessPaymentHandler_Returns400_WhenInputBindingFails(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/process_payment", nil) // No body to cause binding failure
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	paymentCtrl := controller.NewPaymentController(di.NewProcessPayment())

	err := paymentCtrl.ProcessPaymentHandler(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func Test_ProcessPaymentHandler_Returns400_WhenInputContainsInvalidDataTypes(t *testing.T) {
	e := echo.New()
	invalidJSON := `{"amount": "invalid_number", "rideId": 123}` // Assuming amount should be a number and rideId a string
	req := httptest.NewRequest(http.MethodPost, "/process_payment", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	paymentCtrl := controller.NewPaymentController(di.NewProcessPayment())

	err := paymentCtrl.ProcessPaymentHandler(ctx)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
