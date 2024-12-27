package controller

import (
	"encoding/json"

	"github.com.br/gibranct/ride/internal/payment/application/usecase"
	"github.com.br/gibranct/ride/internal/payment/domain/event"
	"github.com.br/gibranct/ride/internal/payment/infra/queue"
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
		return queueController.ProcessPayment.Execute(usecase.ProcessPaymentInput{
			RideId: input.RideId,
			Amount: input.Fare,
		})
	})
	return queueController
}
