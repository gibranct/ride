package controller

import (
	"context"
	"encoding/json"

	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com.br/gibranct/payment/internal/domain/event"
	"github.com.br/gibranct/payment/internal/infra/queue"
)

type QueueController struct {
	ProcessPayment usecase.IProcessPayment
	Queue          queue.Queue
}

func NewQueueController(paymentService usecase.IProcessPayment, queue queue.Queue) *QueueController {
	queueController := &QueueController{
		ProcessPayment: paymentService,
		Queue:          queue,
	}
	queueController.Queue.Consume("rideCompleted.processPayment", func(msg []byte) error {
		var input event.ProcessPaymentEvent
		json.Unmarshal(msg, &input)
		return queueController.ProcessPayment.Execute(context.Background(), usecase.ProcessPaymentInput{
			RideId: input.RideId,
			Amount: input.Fare,
		})
	})
	return queueController
}
