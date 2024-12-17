package controller

import (
	"github.com.br/gibranct/ride/internal/payment/application"
	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentService application.PaymentService
}

func (paymentCtrl *PaymentController) ProcessPaymentHandler(c echo.Context) error {
	return c.String(200, "Payment processed successfully")
}

func NewPaymentController(paymentService *application.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: *paymentService,
	}
}
