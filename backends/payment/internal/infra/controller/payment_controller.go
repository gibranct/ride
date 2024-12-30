package controller

import (
	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	processPayment usecase.IProcessPayment
}

func (paymentCtrl *PaymentController) ProcessPaymentHandler(c echo.Context) error {
	if c.Request().GetBody == nil {
		return c.JSON(400, map[string]string{"message": "Request body is required"})
	}
	var input usecase.ProcessPaymentInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, map[string]string{"message": err.Error()})
	}
	err := paymentCtrl.processPayment.Execute(input)
	if err != nil {
		return c.JSON(500, map[string]string{"message": err.Error()})
	}
	return c.String(200, "Payment processed successfully")
}

func NewPaymentController(processPayment usecase.IProcessPayment) *PaymentController {
	return &PaymentController{
		processPayment: processPayment,
	}
}
