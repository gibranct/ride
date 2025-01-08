package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com.br/gibranct/payment/internal/application/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessPayment(t *testing.T) {
	testCases := []struct {
		name               string
		body               *usecase.ProcessPaymentInput
		expectedStatusCode int
	}{
		{name: "Valid request", body: &usecase.ProcessPaymentInput{Amount: 2323, RideId: uuid.NewString()}, expectedStatusCode: 200},
		{name: "Invalid request body", body: nil, expectedStatusCode: 400},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			jsonBytes, err := json.Marshal(testCase.body)
			assert.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/process_payment", testServer.URL), bytes.NewBuffer(jsonBytes))
			assert.NoError(t, err)
			request.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)

			response, _ := http.DefaultClient.Do(request)
			assert.Equal(t, testCase.expectedStatusCode, response.StatusCode)
		})
	}
}
