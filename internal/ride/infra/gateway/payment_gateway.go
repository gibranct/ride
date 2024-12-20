package gateway

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type PaymentGateway struct{}

func (m *PaymentGateway) ProcessPayment(input ProcessPaymentInput) error {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3002/process_payment", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

type ProcessPaymentInput struct{}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}
