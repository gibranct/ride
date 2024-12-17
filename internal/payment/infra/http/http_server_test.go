package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com.br/gibranct/ride/internal/payment/application"
	"github.com.br/gibranct/ride/internal/payment/infra/controller"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var paymentCtrl = controller.NewPaymentController(application.NewApplication().PaymentService)

func Test_ProcessPayment(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/process_payment", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, paymentCtrl.ProcessPaymentHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, rec.Body.String(), "Payment processed successfully")
	}
}
