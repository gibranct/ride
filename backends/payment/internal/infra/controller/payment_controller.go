package controller

import (
	"context"

	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	processPayment usecase.IProcessPayment
}

func (paymentCtrl *PaymentController) ProcessPaymentHandler(c echo.Context) error {
	var input usecase.ProcessPaymentInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(400, map[string]string{"message": err.Error()})
	}
	if input == (usecase.ProcessPaymentInput{}) {
		return c.JSON(400, map[string]string{"message": "Invalid request body"})
	}
	err := paymentCtrl.processPayment.Execute(context.Background(), input)
	if err != nil {
		return c.JSON(500, map[string]string{"message": err.Error()})
	}
	return c.JSON(200, map[string]string{"message": "Payment processed successfully"})
}

func NewPaymentController(processPayment usecase.IProcessPayment) *PaymentController {
	return &PaymentController{
		processPayment: processPayment,
	}
}
