package controller

import (
	"github.com.br/gibranct/ride/internal/payment/application"
	"github.com.br/gibranct/ride/internal/payment/application/usecase"
	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentService application.PaymentService
}

func (paymentCtrl *PaymentController) ProcessPaymentHandler(c echo.Context) error {
	if c.Request().GetBody == nil {
		return c.JSON(400, map[string]string{"message": "Request body is required"})
	}
	var input usecase.ProcessPaymentInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, map[string]string{"message": err.Error()})
	}
	err := paymentCtrl.paymentService.ProcessPayment.Execute(input)
	if err != nil {
		return c.JSON(500, map[string]string{"message": err.Error()})
	}
	return c.String(200, "Payment processed successfully")
}

func NewPaymentController(paymentService *application.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: *paymentService,
	}
}
